# golang-for-study
Goの勉強用レポジトリ
ビルドとかでファイル名が指定されなくても出来るのは、
恐らくmainパッケージに所属しているものが自動で判別されているのかなと思います

## goのインストール方法（Mac向け）
$ brew install go
→Homebrewを使ってgoをインストール

$ go version
→ go version go1.14 darwin/amd64
　インストールされていることの確認

## ビルド方法
適当な場所にgoのファイル「main.go」とか作成しておく
$ go build
→ディレクトリにディレクトリ名でバイナリが作成される

$ go build -o ファイル名
→指定の名前でbuildファイルを作成できる

## 実行方法
### ビルドしたバイナリを実行
ビルドしたファイルがあるディレクトリまで移動している
$ ./ファイル名
→実行される
### ファイルをそのまま実行
$ go run main.go
→実行される、バイナリは出来ない
