# Golang HandsOn

先日の Golang 勉強会 in Kagawa で行われた [＠qt_luigi](https://twitter.com/qt_luigi) さんの [入門ハンズオン](https://speakerdeck.com/qt_luigi/ru-men-hanzuon) のサンプルコードです。  
こちらを参考に、みなさんもご自分で問題にトライしてみてくださいね。

## タイピングゲームを作ろう！

下記のように実行すると幾つかのキーワードが表示され、それと同じ単語を入力すれば○、間違えば×と表示されるタイピングゲームを作成します。

```
$ go run typing.go
golang
input: golang
○
study
input: study
○
meeting
input: meeting
○
in 
input: in
○
kagawa
input: in
okayama
×
```

では上記について機能拡張をしていきましょう。

問題A. カウントダウン表示
A-1) 最初の単語を表示する前に３秒カウントダウンをしてください。

問題B. 正誤表示
B-1) 最後に「正解数」と「誤り数」を表示してください。
