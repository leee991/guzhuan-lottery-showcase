/**
 * 接口定义 — lottery/tests/static-load.cjs
 * 仅保留模块导出与接口签名。
 */

const fs = require('node:fs');
const http = require('node:http');
const {performance} = require('node:perf_hooks');

function request(route){
  // 内部业务实现略
  throw new Error("not implemented");
}

function stats(rows){
  // 内部业务实现略
  throw new Error("not implemented");
}
