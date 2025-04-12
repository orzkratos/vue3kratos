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
