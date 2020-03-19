/**
 * js 工具类
 */

export default {
  isBlank: v => {
    return v === undefined || v === null || v.length === 0;
  },
  unescapeHTML: str =>
    str.replace(
      /&amp;|&lt;|&gt;|&#39;|&#34;|&quot;/g,
      tag =>
        ({
          '&amp;': '&',
          '&lt;': '<',
          '&gt;': '>',
          '&#39;': "'",
          '&#34;': '"',
          '&quot;': '"',
        }[tag] || tag)
    ),
  uuid: () => {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
      var r = (Math.random() * 16) | 0,
        v = c == 'x' ? r : (r & 0x3) | 0x8;
      return v.toString(16);
    });
  },
  dateFormat: (date, fmt) => {
    var o = {
      'M+': date.getMonth() + 1,
      //月份
      'd+': date.getDate(),
      //日
      'h+': date.getHours(),
      //小时
      'm+': date.getMinutes(),
      //分
      's+': date.getSeconds(),
      //秒
      'q+': Math.floor((date.getMonth() + 3) / 3),
      //季度
      S: date.getMilliseconds(), //毫秒
    };
    if (/(y+)/.test(fmt))
      fmt = fmt.replace(
        RegExp.$1,
        (date.getFullYear() + '').substr(4 - RegExp.$1.length)
      );
    for (var k in o)
      if (new RegExp('(' + k + ')').test(fmt))
        fmt = fmt.replace(
          RegExp.$1,
          RegExp.$1.length == 1
            ? o[k]
            : ('00' + o[k]).substr(('' + o[k]).length)
        );
    return fmt;
  },
};
