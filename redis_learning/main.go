package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
)

const Key = "key"

var luaScript = `
	local key = KEYS[1]
	local result = ARGV[1]
	local combo = redis.call("GET", key) or 0
	combo = tonumber(combo)

	if combo > 0 then
		if result == "win" then
			combo = combo + 1
		else -- lose
			combo = combo - 1
		end
	elseif combo == 0 then
		if result == "win" then
			combo = combo + 1
		else -- lose
			combo = combo - 1
		end
	else -- combo < 0
		if result == "win" then
			combo = 1
		else -- lose
			combo = -1
		end
	end

	redis.call("SET", key, combo)
	return combo
	`

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	userId := "17345"
	gameName := "poker"
	key := fmt.Sprintf("%s_%s", userId, gameName)
	println(key)

	// 初始化
	cmd := rdb.Set(ctx, "playerComboCounter", 0, 0)
	err := cmd.Err()
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(num int) {
			newKey := "playerComboCounter" // 示範用的鍵名，實際使用時請替換成適當的鍵名
			result := "win"                // 或 "lose"，根據遊戲結果設置
			scriptHash, err := rdb.ScriptLoad(ctx, luaScript).Result()
			if err != nil {
				panic(err)
			}

			cmd := rdb.EvalSha(ctx, scriptHash, []string{newKey}, result)
			err = cmd.Err()
			if err != nil {
				panic(err)
			}
			//rdb.Incr(ctx, Key)
			//if err := rdb.Set(ctx, Key, num, 0).Err(); err != nil {
			//	panic(err)
			//}
			wg.Done()
		}(i)
	}

	wg.Wait()

	r, err := rdb.Get(ctx, "playerComboCounter").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("r: ", r)
}
