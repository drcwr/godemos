public static void main(String[] args) {
    // 要验证的字符串
    String str = "service@xsoftlab.net";
    // 邮箱验证规则
    String regEx = "[a-zA-Z_]{1,}[0-9]{0,}@(([a-zA-z0-9]-*){1,}\\.){1,3}[a-zA-z\\-]{1,}";

    
    String str = "status:200 payload:\"3290afeb96aba632aab3d5df13ee738a\"";
    String regExstatus = "status:[0-9]{1,3}";
    Pattern patternstatus = Pattern.compile(regExstatus);
    Matcher matcherstatus = patternstatus.matcher(str);
    if (matcherstatus.find()) {
        String strstatus = matcherstatus.group(1);
    }

    String regExpayload = "payload:\"(.*?[^\\\\])\"";
    Pattern patternpayload = Pattern.compile(regExpayload);
    Matcher matcherpayload = patternpayload.matcher(str);
    if (matcherpayload.find()) {
        String strpayload = matcherpayload.group(1);
    }
}
