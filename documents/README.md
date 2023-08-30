```
yarn init
yarn add -D vuepress
yarn docs:dev
```
```
set NODE_OPTIONS=--openssl-legacy-provider
```
```
{
  "scripts": {
    "docs:dev": "vuepress dev docs",
    "docs:build": "vuepress build docs"
  }
}
```
