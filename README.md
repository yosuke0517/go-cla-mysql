# go-cla-mysql

# フレームワーク
- echo: 使ったことあるから程度

# ミドルウェア
- DB: MySQL(5.7)
  - 8系とgoのライブラリの対応状況が不明確なので一旦5.7を使う
  - ドライバー: "github.com/go-sql-driver/mysql"を使用
- セッション管理: Redis

# 環境変数
- "github.com/joho/godotenv"を使用

# 環境構築
- https://fuzzy-hunter-3bf.notion.site/golang-docker-afad7bfb494740b3910834a489452784

# フレームワーク
- echo
  - ルーティング（Get, Postの分け方とか）が簡単だったけど、それでしか使ってない

# ディレクトリ構成
- TODO notionにまとめて公開する

# 決め事
- inputPortで宣言する関数名
- FindAll() (domain.Todosなどの複数形, error)
- FindByID(id int) (domain.Todosなどの複数形, error)
- Save(domain.Todoなどの単数系) (int64, error)
- DeleteByID(int) error

- inputPortで宣言した関数は/usecases/interactorに同名のファイルを作成し実装する

# handler
- フレームワークなどを使うとcontroller内にルーティングとメソッドのマッピングを書けたりするが今回はhandlerの中でルーティングとメソッドのマッピングを行う


# 確認したいこと
- 依存関係の注入（inject）とルーティングの登録はどの階層？？（現状トップレベル）
- レスポンスの共通部品はpresenterでOK??
- httpメソッドのヘルパーをroutesに作ったけどここにあるの違和感ある？？
- TodoInputPortの返り値ポインタにしているが値でもいいか？？