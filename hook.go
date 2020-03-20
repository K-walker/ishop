public function exec_shell() {
  set_time_limit(3 * 60); //最大过期时间3分钟
  $shellPath = "/home/www/ishop";
  $cmd = "cd $shellPath && sudo git pull && sh CI.sh";
  $res = $this -> doShell($cmd);
  print_r($res); // 主要打印结果给github记录查看，自己测试时查看
}
