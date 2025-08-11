# Makefile

# 目标：更新所有依赖到符合 package.json 中的版本范围
update:
	npm ls
	npm fund
	npm audit
	npm audit fix
	npm update
	npx npm-check-updates -u
	npm install

# 读取 tsconfig.json 按照配置检查类型和语法是否正确
check:
	npx tsc --noEmit
