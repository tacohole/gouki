# !/bin/sh

go build -o gouki
cp .env.sample .env
# append to $PATH in user's rc