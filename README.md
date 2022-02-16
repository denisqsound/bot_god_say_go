# bot_god_say_go

go build -o bin/bot_god_say_go -v .
scp -r /Users/dekhakhalkin/Desktop/BIBLE_APP/ root@5.63.158.231:/root/bible_app
tmux new-session -t bot
tmux attach -t bot
Ctrl+ b d or Ctrl+ b :detach


heroku ps:scale clock=0