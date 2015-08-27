# Go 言語の練習用 repository

Go 言語の練習用 repository です。

## 開発環境

この repository は [gb] を使った開発環境です。

```shell
C:>go get -v github.com/constabulary/gb/...
github.com/constabulary/gb (download)
github.com/constabulary/gb
github.com/constabulary/gb/cmd
github.com/constabulary/gb/vendor
github.com/constabulary/gb/cmd/gb
github.com/constabulary/gb/cmd/gb-vendor
```

`go get` コマンドを使って導入するのではなく，普通に `git clone` して環境を作って下さい。ビルドには `gb build` コマンドを使います。

```shell
C:\workspace>git clone https://github.com/spiegel-im-spiegel/go-practice.git
Cloning into 'go-practice'...
remote: Counting objects: 18, done.
remote: Compressing objects: 100% (15/15), done.
remote: Total 18 (delta 1), reused 14 (delta 1), pack-reused 0
Unpacking objects: 100% (18/18), done.
Checking connectivity... done.

C:\workspace>cd go-practice

C:\workspace\go-practice>gb build
FATAL command "build" failed: failed to resolve import path "fileconv": cannot find package "golang.org/x/text/encoding/japanese" in any of:
        C:\Go\src\golang.org\x\text\encoding\japanese (from $GOROOT)
        C:\workspace\go-practice\src\golang.org\x\text\encoding\japanese (from $GOPATH)
        C:\workspace\go-practice\vendor\src\golang.org\x\text\encoding\japanese
```

外部パッケージがない場合には `gb vendor fetch` コマンドで導入します。

```shell
C:\workspace\go-practice>gb vendor fetch golang.org/x/text/encoding/japanese
fetching recursive dependency golang.org/x/text/encoding
fetching recursive dependency golang.org/x/text/transform
```

では，最後までビルドしてみます。

```shell
C:\workspace\go-practice>gb build
golang.org/x/text/transform
golang.org/x/text/encoding/internal/identifier
golang.org/x/text/encoding
golang.org/x/text/encoding/internal
golang.org/x/text/encoding/japanese
fileconv

C:\workspace\go-practice>bin\fileconv.exe -version
fileconv version 0.1.0

C:\workspace\go-practice>tree .
C:\WORKSPACE\GO-PRACTICE
├─bin
├─pkg
│  └─windows
│      └─amd64
│          └─golang.org
│              └─x
│                  └─text
│                      └─encoding
│                          └─internal
├─src
│  └─fileconv
└─vendor
    └─src
        └─golang.org
            └─x
                └─text
                    ├─encoding
                    │  ├─charmap
                    │  ├─htmlindex
                    │  ├─ianaindex
                    │  ├─internal
                    │  │  └─identifier
                    │  ├─japanese
                    │  ├─korean
                    │  ├─simplifiedchinese
                    │  ├─testdata
                    │  ├─traditionalchinese
                    │  └─unicode
                    └─transform
```

## ライセンス

再利用しようという奇特な方がおられるなら CC0 ライセンスでどうぞ。

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](http://creativecommons.org/publicdomain/zero/1.0/deed.ja)

[gb]: http://getgb.io/ "gb - A project based build tool for Go"
