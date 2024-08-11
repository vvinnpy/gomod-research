# readme
## 版本和模块
- require github.com/forifido/gomod-research v1.0.0     
  > 正确
- require github.com/forifido/gomod-research v2.0.0     
  > 错误, 使用 require github.com/forifido/gomod-research v1.0.1-0.20240811134202-0ebcf509e418
- require github.com/forifido/gomod-research/v2 v2.1.0
  > 正确
- require github.com/forifido/gomod-research/v2 v2.2.0
  > 正确
- require github.com/forifido/gomod-research/v2 v2.3.0
  > 无法导入, 根目录和 v2 子目录定义了相同 mod

