package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-cla-practice/routes"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample swagger server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /api/v1
func main() {
	loadEnv()
	// ルーティング(handler登録)
	routes.InitRouting()
	fmt.Println("start")
}

func loadEnv() {
	err := godotenv.Load(".env")
	//もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("環境変数ファイルを読み込み出来ませんでした: %v", err)
	}
}
