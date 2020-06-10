# A sample to generate a C shared library using cgo

Go の cgo を使って共有ライブラリを作る方法を簡潔にまとめたプロジェクトです。

## このプロジェクトでできること

[ojichat](https://github.com/greymd/ojichat) の機能の一部をライブラリ化し、C言語で利用します。

## ライブラリ生成

プロジェクトのルートにて下記を実行。  
`_c` の中にライブラリとヘッダファイルが生成される。

出力するライブラリの拡張子は環境に合わせること。

```shell script
$ go build --buildmode=c-shared -o _c/libojichat.dll ojichat.go
```

## ライブラリを利用するC言語のコードから実行ファイルを生成

ライブラリや出力する実行ファイルの拡張子は環境に合わせること。

```shell script
$ cd _c
$ gcc libojichat.dll -o ojichat.exe ojichat.c
```

## 実行ファイルの実行例

```shell script
$ ./ojichat.exe
お名前： かぼ
かぼちゃんのお目々、キラキラ😊してるネ、😚😂ホント可愛すぎだ、よ、〜😃☀ 😍マッタクも、う、（笑）💗😃♥
```

## メモ

* Go の関数は複数の値を返せるが、C では構造体になるっぽい
    * チャットメッセージとエラーを一緒に返すと扱いにくそうなのでやめた
    * 代わりに、エラーの場合は空文字を返している

## 関連リポジトリ

* [dart-ffi_cgo_ojichat](https://github.com/kaboc/dart-ffi_cgo_ojichat)
    * 作ったライブラリを Dart で利用するプロジェクト
