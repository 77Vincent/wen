---
title: AWS IAM 认证与授权
date: 2024-11-26T02:01:58+05:30
tags: [ computer-science, aws, iam, authentication, authorization ]
categories: study
canonicalUrl: https://wenstudy.com/posts/aws-authentication-and-authorization/
---
 
AWS IAM 是个极其重要却不甚被人熟悉的服务，全称是 **Identity and Access Management**，即身份和访问管理。用于做认证和授权。本文总结了关于一切关于 IAM 重要的知识点。

<!-- more -->

## 访问控制 (Access Control)
你的服务即不能无保护，让任何人都可以访问，也不能完全封闭以至于用户和管理者也无法使用，因此需要认证和授权机制。

### 认证
认证 (`Authentication`) 即证明你是自己人，通过用户名密码、证书、指纹等方式。

> 失败的认证 = 401错误

### 授权
授权 (`Authorization`) 即在知道你自己人后，证明你有权访问目标资源，通过角色、权限等方式。

> 失败的授权 = 403错误

## 身份 (IAM Identity)

### 用户
用户（IAM User）是 AWS 账户下的一个实体 (`Entity` / `Principal`)，可以登录 AWS 控制台，也可以通过 API 访问 AWS 服务，_理解为一个法人就好_；其可以属于多个组 (`Groups`)，也可以拥有多个角色 (`Roles`)。


#### 根用户

每个 AWS 账户都有一个**根用户** (`Root User`)，拥有所有权限，仅靠用户名密码保护，一旦泄露，后果严重。所以，有一套最佳实践：
1. 不在根用户下创建访问密钥 (`Access keys`)。
2. 使用很长的密码，存在安全的密码管理器中，并定期更换。
3. 启用多因素身份验证 (`MFA`)。
4. 平常尽量别根用户，而创建拥有 `AdministratorAccess` 权限的 IAM 用户。

> 即便是 `AdministratorAccess` 权限，也不是无限制的。

#### 访问密钥
**访问密钥**（`Access Keys`）为程序或命令行工具提供访问权限，类似用户名密码，一份 `access key` 由 `ID` 和 `Secret` 组成，`ID` 用于标识。

```bash
$ cat ~/.aws/configure
aws_access_key_id = AKIA2VL4IAFXXXXXXXXX
aws_secret_access_key = abcdefghijklmnopqrstuvwxyz1234567890 
```

这么做的好处是，可以颁发诸多份密钥，而不用共享密码，并可以随时撤销。_归根结底，给人使用的用户名密码不适合程序使用。_

#### 密钥轮换
**密钥轮换** (`Key Rotation`) 是一种安全最佳实践，以减少泄露风险。对于被 EC2 使用的的 IAM 角色（`roles`），密钥天然自动轮换。
但如果你的密钥是给自己的程序用的，就要手动实现轮换。基本步骤是：

1. 创建新密钥。
2. 更新程序配置。
3. 测试新密钥。
4. 删除旧密钥。


### 组

### 角色

### 身份提供商

### 身份联合

## IAM 政策 (IAM Policies) 
IAM 政策是 JSON 格式的文档，用于指定 _谁能在什么条件下使用什么资源_。IAM 提供了一些内置的策略，也可以自定义。几个要点：

1. 一个 IAM 政策可以附加到无限多的用户、组、角色身份上。
2. 而一个身份最多可以被贴上 10 个 IAM 政策 (每个政策不超过 6,144 字符)。
3. 如果政策间冲突，拒绝 (Explicit deny) 优先级最高。

![IAM user-group-role](/images/aws-authentication-and-authorization/iam-user-group-role.png "IAM user-group-role")

以下政策的作用是允许来自 `111.111.111.111` 的 IP 地址访问 S3 服务。
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:*",
      "Resource": "*",
      "Condition": {
        "IpAddress": {
          "aws:SourceIp": "111.111.111.111"
        }
      }
    }
  ]
}
```
