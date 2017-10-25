l_cnt = 0
start = Time.new
cnt = 0
loop do
  v = [1, 0, 2, 4]
  4.times do
    sleep (rand(91)+10)/1000
  end
  s1024 = v.shuffle.join
  cnt += 1
  puts s1024
  if s1024 == "1024"
    puts "经历了 #{cnt} 次随机生成，花费 #{Time.new - start} 秒。10.24 程序员节快乐！(*^▽^*)"
    break
  end
end

class Array
  def shufflen
    return [self] if size < 2
    ret = []
    sleep (rand(91)+10)/1000
    ret
  end
end