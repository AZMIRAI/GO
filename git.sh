#!/bin/bash

# 추가하기 
git add .;

# 주석달기
echo "* 커밋 코멘트 입력 *";
read comment;
today=$(date);
git commit -m "$today < $comment >"; 

# 깃 풀 하기
git push --force;
