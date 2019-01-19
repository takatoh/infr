# infr
INtegral in FRequency domain, written in golang.

等時間間隔の時刻歴波形を周波数領域で積分します。

時間領域の時刻歴波形をフーリエ変換を用いて周波数領域に変換し、積分したのち、
フーリエ逆変換で時間領域に戻す、という利用法を想定しています。

## Install
``` go get github.com/takatoh/infr```

## Usage
infr.Integrate 関数を使います。

``` y := infr.Integrate(x, n, dt, v0)```

x がフーリエ変換されたデータ（複素数）、n がデータ数、dt が時間間隔、v0 が初期値です。

フーリエ変換/逆変換には [github.com/takatoh/fft](https://github.com/takatoh/fft)
ライブラリが利用できます。

## License
MIT License
