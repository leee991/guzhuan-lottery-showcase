/**
 * 接口定义 — lottery/lib/static-cache.js
 * 仅保留模块导出与接口签名。
 */


const fs = require('node:fs/promises');
const crypto = require('node:crypto');
const zlib = require('node:zlib');
const {promisify} = require('node:util');
