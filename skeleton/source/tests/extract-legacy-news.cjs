/**
 * 接口定义 — source/tests/extract-legacy-news.cjs
 * 仅保留模块导出与接口签名。
 */

const fs = require('node:fs/promises');
const path = require('node:path');
const {chromium} = require(process.env.CMS_PLAYWRIGHT_MODULE || '/path/to/local');
