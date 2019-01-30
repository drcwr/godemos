Map<String, String> map = new HashMap<String, String>();
String test = "pet:cat::car:honda::location:Japan::food:sushi";
String[] test1 = test.split("::");

for (String s : test1) {
    String[] t = s.split(":");
    map.put(t[0], t[1]);
}

for (String s : map.keySet()) {
    System.out.println(s + " is " + map.get(s));
}