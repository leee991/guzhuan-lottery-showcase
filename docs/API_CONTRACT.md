# 接口与命令清单 — 开奖站 + Go CMS + Bot

## 1. Go CMS 接口（前缀 `/api/cms`）

| Method | Path | 权限 | 说明 |
|---|---|---|---|
| `GET` | `/health` | 公开 | 健康检查（磁盘满或备份超时返回通用 503，不暴露内部路径） |
| `GET` | `/public/content?kind=news&slug=` | 公开 | 已发布且到定时发布时间的内容，支持 ETag |
| `GET/HEAD` | `/public/article/{slug}` | 公开 | 服务端渲染文章页（无 JS、严格 CSP、规范 Canonical） |
| `GET` | `/public/sitemap.xml` | 公开 | 已发布资讯的 XML Sitemap |
| `GET/HEAD` | `/public/media/{id}.{jpg\|png}` | 公开 | 128-bit 随机 ID 图片，长效不可变缓存 |
| `GET` | `/admin/me` | 登录 | 当前管理员角色与读写权限 |
| `GET/POST/PUT` | `/admin/content` · `/{id}` | 编辑 | 列表、新建、更新（带 `version` 乐观锁，冲突 409） |
| `GET/POST` | `/admin/content/{id}/revisions` · `/restore` | 编辑 | 历史版本查看与非破坏性恢复 |
| `GET/POST` | `/admin/media` · `/{id}/archive` · `/restore` | 编辑 | 上传（<=5MiB）、归档、恢复 |
| `GET/POST/PUT` | `/admin/categories` · `/admin/tags` | 编辑 | 最多 4 级分类树（100 分类 / 300 标签，无环校验） |
| `GET/PUT/DELETE` | `/admin/drafts/{key}` | 编辑本人 | 按 `Actor.ID` 隔离的私有自动保存草稿（7 天过期） |
| `GET/POST` | `/admin/backups` · `/{id}/verify` · `/restore` | 超管 | 建快照、SHA-256 校验、带安全快照的恢复 |
| `GET` | `/admin/metrics/history?range=day\|week` | 登录 | 7 天分钟级请求数、错误率、P95、堆内存和宿主机指标 |
| `GET/POST` | `/admin/inspection` · `/prepublish` | 登录 | 全站只读巡检和发布前预检 |
| `GET` | `/admin/audit` · `/admin/backup-status` | 登录 | 最近 200 条操作审计、只读外部备份状态 |

## 2. Node 主站核心接口（`lottery`）
- 登录、SVG 图形验证码、后台 RBAC 安全策略
- 开奖查询（秒级缓存）、下单（`Idempotency-Key` 幂等 + 股数范围校验）、结束前 10 秒编辑锁定
- 余额/流水查询、单笔上限 50,000 提现校验

## 3. Telegram 机器人命令
- `/history`：查最近 10 期（UTC）
- `/issue <期号>`：按期号精确查
- `/stats`：最近 100 期的和值和单双次数统计
- `/status`：看数据源同步是否正常
