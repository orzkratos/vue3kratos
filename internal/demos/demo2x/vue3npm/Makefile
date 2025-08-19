# Makefile

# WARNING: Force updates all dependencies to the latest versions specified in package.json.
# This can be a destructive action and might introduce breaking changes.
update:
	npm ls
	npm fund
	npm audit
	npm audit fix
	npm update
	npx npm-check-updates -u
	npm install

# Performs a static type check of the project using the TypeScript compiler.
check:
	npx tsc --noEmit
