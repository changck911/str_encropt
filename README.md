# AES 加密/解密工具

## .env 文件配置
METHOD=(1 encrypt 2 decrypt)<br>
AES_KEY_STR=1234567 (至少七個字)<br>
STR=example<br>

## 使用說明
1. 加密：設置 METHOD=1，STR 為要加密的文本
2. 解密：設置 METHOD=2，STR 為加密後的結果（包含密文和IV，格式為 "密文:IV"）

注意：加密時會生成隨機的 IV，並與密文一起輸出。解密時需要輸入完整的加密結果（包括IV）。