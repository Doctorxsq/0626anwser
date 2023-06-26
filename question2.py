#! /usr/bin/env python
# -*- coding: utf-8 -*-
from tutorial import person_pb2
 

pers = person_pb2.AddressBook()
p1 = pers.people.add()
p1.id = 1
p1.name = 'snake'
p1.email = 'abc@qq.com'
p2 = pers.people.add()
p2.id = 2
p2.name = 'qiezi'
p2.email = '123@qq.com'
 
# 对数据进行序列化
data = pers.SerializeToString()
 
# 对已经序列化的数据进行反序列化
target = person_pb2.AddressBook()
target.ParseFromString(data)
print(target.people[1].name)  #  打印第一个 person name 的值进行反序列化验证
