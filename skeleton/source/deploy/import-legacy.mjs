/**
 * 接口定义 — source/deploy/import-legacy.mjs
 * 仅保留模块导出与接口签名。
 */


import fs from 'node:fs/promises';

async function api(route,method='GET',body){
  // 内部业务实现略
  throw new Error("not implemented");
}

// Discovered route endpoints in this module (read-only contract summary):
//   - /api/admin/login
//   - /api/admin/ops/logout
