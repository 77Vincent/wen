const hash = md5.create()
hash.update(location.pathname)
const id = hash.hex()

const gitalk = new Gitalk({
    clientID: 'Ov23lit0mk6PdYjm9mYF',
    clientSecret: '9bb9c80a13eee33cc6135079470ad110c5d9927e',
    repo: 'wen-blog-comments',      // The repository of store comments,
    owner: '77Vincent',
    admin: ['77Vincent'],
    id,      // Ensure uniqueness and length less than 50
})

gitalk.render('gitalk-container')
