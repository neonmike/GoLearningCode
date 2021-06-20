//
//Copyright [yyyy] [name of copyright owner]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**@file Test-Closure.go
* @brief  学习Go中关键内容的用法
* @details  Go的单元测试
* @author      要你命三千又三千  any question please send mail to liukeshanren@gmail.com ,www.lksr.net
* @date        2018-8-17
* @version     V1.0
* @attention
* 硬件平台: windows 10 家庭版
* SDK版本：1.16.5 windows amd64
*/
package Mashal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
}
func (this *Monster) Store() bool{
	data ,err:= json.Marshal(this)
	if err != nil {
		fmt.Println("Marshal err=",err)
		return false
	}
	filePath:="d:/monster.ser"
	err=ioutil.WriteFile(filePath,data,066)
	if err != nil {
		fmt.Println("Write file err",err)
		return false
	}
	return true
}
func (this*Monster) ReStore() bool{
	filePath:= "d:/monster.ser"
	data,err :=ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile errr=",err)
		return  false
	}
	err= json.Unmarshal(data,this)
	if err!=nil {
		fmt.Println("UnMarshal err=",err)
		return false
	}
	return true
}
