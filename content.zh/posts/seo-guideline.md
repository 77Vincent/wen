---
title: 网站的搜索引擎优化
date: 2024-11-14T02:01:58+05:30
description: 使用HTTPS，每个页面都有H1, H2，优化网站速度，适配移动端，优化标题和描述，使用友好的URL，站点地图（SiteMap），Robots.txt，使用Schema标记等措施可以提升网站在搜索引擎中的排名。
categories: study
tags: [ computer-science, marketing ]
canonicalUrl: https://wenstudy.com/posts/seo-guideline/
---

<!-- more -->
主要有以下措施，排名不分先后:

## 使用HTTPS
HTTPS的加密使得用户和网站敏感信息不会在传输过程中泄露，提升用户信任度，增加访问可能性。搜索引擎会优先显示HTTPS网站。

```
https://wenstudy.com/
```

## 每个页面都有H1, H2
搜索引擎会优先抓取H1和H2等标签中的信息，通过这些标签识别页面的主要内容和重要信息，利于排名提升。缺乏 `h1` 与 `h2` 的页面会被搜索引擎认为是不完整的页面，降低排名。
```html
<h1>这是标题</h1>
<h2>这是副标题</h2>
```

## 优化网站速度
网页加载速度影响排名，因此建议
1. 用内容分发网络（CDN）加速资源加载。
2. 压缩资源。
3. 控制图片分辨率。
4. 使用迷你版本的代码比如 `foo.min.js`, `foo.min.css`。 

## 适配移动端
它可以提升移动端用户体验，优化谷歌搜索排名。 根据Google的移动优先索引策略，一个对移动设备友好的网站将在搜索结果中获得更高的排名。

## 优化标题和描述
标题（title）和描述（description）是HTML的header中的两个重要元素。其内容应当简洁明了，能够准确描述页面的内容，吸引用户点击，并进而提升排名。
```html
<head>
    <meta name="description" content="这是描述">
    <title>这是标题</title>
</head>
```

## 使用友好的URL
URL应当简洁明了，自然反应页面内容，便于用户理解和记忆。当页面里的URL包含关键词时，搜索引擎会更容易识别页面内容，提升排名。
```
https://wenstudy.com/posts/seo-guideline/
```

## 站点地图（SiteMap）
站点地图是一个XML文件，包含了网站的所有页面链接，可以帮助搜索引擎更好地抓取网站内容。
一般来说，`sitemap` 的URL是`https://example.com/sitemap.xml`。

```xml
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>https://wenstudy.com/</loc>
        <lastmod>2024-11-14</lastmod>
        <changefreq>daily</changefreq>
        <priority>0.8</priority>
    </url>
</urlset>
```

## Robots.txt
`robots.txt` 是一个文本文件，用于告诉搜索引擎哪些页面可以抓取，哪些页面不可以抓取。一般来说，`robots.txt` 的URL是`https://example.com/robots.txt`。

```txt
User-agent: *
Disallow: /admin/
Disallow: /private/
```

## 使用Schema标记
Schema标记是一种结构化数据标记，可以帮助搜索引擎更好地理解网站内容，提升搜索结果的展示效果。 将以下script放在页面的`<head>`标签中。

```html
<script type="application/ld+json">
{
    "@context": "https://schema.org",
    "@type": "Article",
    "headline": "这是标题",
    "datePublished": "2024-11-14",
    "image": "https://wenstudy.com/image.jpg",
    "author": {
        "@type": "Person",
        "name": "作者"
    },
    "publisher": {
        "@type": "Organization",
        "name": "网站名称",
        "logo": {
            "@type": "ImageObject",
            "url": "https://wenstudy.com/logo.jpg"
        }
    }
}
</script>
```
