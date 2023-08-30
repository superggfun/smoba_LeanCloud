module.exports = {
  head: [
    [
      ['link', { rel: 'icon', href: 'ay.ico' }],
    ]
  ],
  title: 'Smoba',  // 设置网站标题
  description: 'Smoba',
  lang: 'zh-CN',
  themeConfig: {
    nav: [
    ],
    sidebar: [
      {
        title: '指南',   // 必要的
        collapsable: true, // 可选的, 默认值是 true,
        sidebarDepth: 0,    // 可选的, 默认值是 1
        children: [
          '/'
        ]
      },
      {
        title: '快速开始',
        collapsable: true, // 可选的, 默认值是 true,
        sidebarDepth: 0,    // 可选的, 默认值是 1
        children: [
          '/start/config',
          '/start/leancloud',
          '/start/serverless',
          '/start/wxpush',
        ]
      },
      {
        title: '高级',
        collapsable: true, // 可选的, 默认值是 true,
        sidebarDepth: 0,    // 可选的, 默认值是 1
        children: [
          '/senior/markdown',
          '/senior/compile',
        ]
      }
    ]
  }
}