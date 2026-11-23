/**
 * 接口定义 — lottery/ops/export-snapshot.cjs
 * 仅保留模块导出与接口签名。
 */

const fs=require('node:fs');
const path=require('node:path');
const crypto=require('node:crypto');
const {readSnapshot}=require('../lib/sqlite-store');
