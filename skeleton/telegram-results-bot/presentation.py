"""
接口定义 — telegram-results-bot/presentation.py
仅保留模块结构与导出签名。
"""

import datetime as dt
from html import escape

def _values(row):
    # 内部业务实现略
    raise NotImplementedError('not implemented')

def classify_sum(total):
    # 内部业务实现略
    raise NotImplementedError('not implemented')

def _shape(values):
    # 内部业务实现略
    raise NotImplementedError('not implemented')

def _draw_line(row, latest=False):
    # 内部业务实现略
    raise NotImplementedError('not implemented')

def _local_time(value):
    # 内部业务实现略
    raise NotImplementedError('not implemented')

def channel_caption(rows, source='public-feed.example.com'):
    # 内部业务实现略
    raise NotImplementedError('not implemented')

def _events(values):
    # 内部业务实现略
    raise NotImplementedError('not implemented')

def _alert_text(header, values, footer):
    # 内部业务实现略
    raise NotImplementedError('not implemented')

def _event_keys():
    # 内部业务实现略
    raise NotImplementedError('not implemented')

def stats_alert(rows):
    # 内部业务实现略
    raise NotImplementedError('not implemented')

def gaps_alert(rows):
    # 内部业务实现略
    raise NotImplementedError('not implemented')
