package routes

import (
	"fmt"
	"go-cla-mysql/adapters/controllers"
	"go-cla-mysql/injector"
	"net/http"
)

func InitRoutign() {
	todo := injector.InjectTodo()
	http.HandleFunc("/health", get(controllers.HealthCheckController()))
	http.HandleFunc("/hello", get(controllers.HelloController()))
	// memo: スラッシュの有無が厳格なのでQuery取るときは末尾スラッシュ入れないと入ってこない
	// TODO パラメータがないやつも入ってくるようにできないかな？？
	http.HandleFunc("/api/v1/todo/get/", get(todo.GetAll()))
	http.HandleFunc("/api/v1/todo/getOne/", get(todo.GetOne()))
	http.HandleFunc("/api/v1/todo/create/", post(todo.Create()))
	// memo: ListenAndServeはHandleFuncの登録後
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil))
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
