/**
 * 接口定义 — lottery/lib/sqlite-store.js
 * 仅保留模块导出与接口签名。
 */

const fs = require('node:fs');
const path = require('node:path');
const {DatabaseSync} = require('node:sqlite');

function readDatabase(db) {
  // 内部业务实现略
  throw new Error("not implemented");
}

function createStore(jsonFile) {
  // 内部业务实现略
  throw new Error("not implemented");
}

function readSnapshot(jsonFile, {engine} = {}) {
  // 内部业务实现略
  throw new Error("not implemented");
}
