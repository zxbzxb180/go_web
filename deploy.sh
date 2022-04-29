kill 9 $(pgrep webserver)

cd ~/go_web/
git pull https://github.com/zxbzxb180/go_web.git
cd webserver/
./webserver &