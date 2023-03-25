# 生成静态文件
npm run docs:build

# 进入生成的文件夹
cd docs/.vuepress/dist

# 如果是发布到自定义域名
echo 'docs.bbs-go.com' > CNAME

git init
git add -A
git config user.name 'bbs-go'
git config user.email 'bbs-go@qq.com'
git commit -m 'deploy docs'

git push -f git@github.com:bbs-go/bbs-go-docs.git master
