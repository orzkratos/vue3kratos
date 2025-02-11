# Makefile

# 目标：更新所有依赖到符合 package.json 中的版本范围
update:
	#nvm ls
	npm ls
	npm fund
	npm audit
	npm audit fix
	npm update

# 目标：更新所有依赖到最新版本
upgrade:
	#nvm install node
	#nvm use node
	npx npm-check-updates -u
	npm install
