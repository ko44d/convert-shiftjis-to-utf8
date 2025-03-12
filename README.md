# convert-shiftjis-to-utf8

このリポジトリは、Goで実装されたシンプルなツールです。Shift-JISでエンコードされたCSVファイルをUTF-8に変換します。

## 特徴

- Shift-JISエンコードのCSVファイルをUTF-8に変換します。
- Goの標準ライブラリのCSVパッケージと、`golang.org/x/text/encoding/japanese` パッケージを利用して、正確にShift-JISのデコードを行います。
- コマンドラインから簡単に実行でき、エラー発生時には明確なメッセージで通知します。

## 使い方

以下のコマンドをターミナルで実行し、CSVファイルを変換してください:

```bash
go run main.go <入力ファイル> <出力ファイル>
```

ここで、`<入力ファイル>` にはShift-JISでエンコードされたCSVファイルのパスを、`<出力ファイル>` には変換後のUTF-8ファイルのパスを指定します。

## 依存パッケージの更新手順

このリポジトリはGo Modulesを利用して依存パッケージを管理しています。最新の状態に更新するための手順は以下の通りです:

1. ターミナルを開き、リポジトリのルートディレクトリに移動します。
2. 以下のコマンドを実行して、すべての依存パッケージを最新バージョンに更新します:

   ```bash
go get -u ./...
   ```

3. 不要な依存パッケージを削除し、`go.mod` と `go.sum` ファイルを整理するために、以下のコマンドを実行します:

   ```bash
go mod tidy
   ```

4. これで依存パッケージの更新は完了です。
