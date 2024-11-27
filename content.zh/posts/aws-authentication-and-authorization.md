---
title: AWS IAM 认证与授权
date: 2024-11-25T02:01:58+05:30
tags: [ computer-science, aws, iam, authentication, authorization ]
categories: study
canonicalUrl: https://wenstudy.com/posts/aws-authentication-and-authorization/
---
 
AWS IAM 是个极其重要却不甚被人熟悉的服务，全称是 Identity and Access Management，即身份和访问管理。用于做认证和授权。

<!-- more -->

## 认证与授权
你的服务即不能无保护，让任何人都可以访问，也不能完全封闭以至于用户和管理者也无法使用，因此需要认证和授权机制。

### 认证 (Authentication)


### 授权 (Authorization)

## 身份 (IAM Identity)

### 根用户 (Root User)

每个 AWS 账户都有一个根用户，拥有所有权限，仅用用户名密码保护。不推荐经常使用根用户，因为它是上帝，一旦泄露，后果严重。

## 政策 (IAM Policies) 
