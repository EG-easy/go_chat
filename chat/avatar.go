package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

//ErrNoAvatarはAvatarインスタンスがアバターのURLを返すことができない場合に発生するエラーです
var ErrNoAvatarURL = errors.New("chat:アバターのURLを取得できません。")
//Avatarはユーザーのプロフィール画像を表す型です
type Avatar interface{
	//GetAvatarURLは、指定されたクライアントのアバターのURLを返します。
	//問題が発生した場合にはエラーを返します。特にURLを取得することができなかった場合は、ErrNoAvatarURLを返します。
	GetAvatarURL(ChatUser)(string, error)
}

type TryAvatars []Avatar
func (a TryAvatars)GetAvatarURL(u ChatUser)(string, error){
	for _, avatar := range a {
		if url , err := avatar.GetAvatarURL(u); err ==nil{
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

type AuthAvatar struct {}

var UseAuthAvatar AuthAvatar
func(_ AuthAvatar)GetAvatarURL(u ChatUser)(string, error) {
	url := u.AvatarURL()
	if url != ""{
		return url, nil
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct {}
var UseGravatar GravatarAvatar
func(_ GravatarAvatar)GetAvatarURL(u ChatUser)(string, error) {
	return "//www.gravatar.com/avatar/" +u.UniqueID() , nil
}

type FileSystemAvatar struct{}
var UseFileSystemAvatar FileSystemAvatar
func (_ FileSystemAvatar)GetAvatarURL(u ChatUser)(string, error){
	if files, err := ioutil.ReadDir("avatars"); err ==nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}