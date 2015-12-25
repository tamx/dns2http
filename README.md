# dns2http

このツールは DNS の TCP パケットを HTTP でリダイレクトする
クライアントと CGI のセットです。

使用法は servercgi にお使いの DNS サーバの IP を書き込み、
Web サーバ側の CGI として仕込んで下さい。
その後、利用したいマシンから
% dns2http [sport] [serverurl]
とすることで、 sport で DNS パケットを TCP で待ち受けるようになります。

プロキシを使用したい場合には、環境変数で設定して下さい。
DNS の TCP しか対応していません。
UDP を TCP に変換するのは別のツールを利用して下さい。
