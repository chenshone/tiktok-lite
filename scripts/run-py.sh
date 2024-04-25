#!/usr/bin/env bash

echo "Please Run Me on the root dir, not in scripts dir."

src="src"

# 创建数据库表 (把src/model/models.py文件中的main函数注释去掉)

cd "$src"

python -m model.models

