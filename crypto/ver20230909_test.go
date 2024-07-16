package crypto

import (
	"testing"
)

func TestVer20230909Lua(t *testing.T) {
	enText := `@d3yIKW1tYpJXe9xZm/Q01wYzQLohl5ixRZKTF22v7q65115K45P2hocJ/JEJntSUPIiSNLXvj9tmdDgUx++Gzmqw//sWcEqSH57G4qLf7msyiGMrpc2e4Roa9H/P0zcgW/Kc4n7xKapBcdm+xznDlus/URgGqEF1kLRlKtm85hJNTfwpWT+unLYC/C25shf52jewF/jI1qiXPVcze1ecQrY1HNad+PBY8UrreRuHU7rsslZ/n1NjCcexJ3SjjNEdNtrKVHEd9RwbYIr8d5vLyZNRCj81u3x7oIL1vw4eSasJ8t6927XvT5mZnvU9/3r6VDYt1ZNEpkHdfnxBlD+PAVos5N3c1Y0UZvHYI67DXh5Sjy6KsfXEXuJPfHUuYPBMquhRqfFZ1Nw4d3CNfxBalgS5XFN4pHLZx63B4IYLhaqKWjzIkUSQjf4JzsS5FwFWIsvVNvl9r7isVxNAHwWdFOAC5O69KBPZBQrt4Kip7T3p41fExn/mnJh2TuRjcYaxkxqvmDaGyxnoxo3MQFaYoKKck1wlll1NwiGImhTnetun+8r5JrJWBOEelFrl0yO3dEOB4AZlMpbbUydZVqgz8XfFQ9yqpJMiK2MqSqIXd7hA2gWPOM5fu9TJhz4sXNdr0X3B6+KkAzxe+aTld+kbRuhUsGJM/DlBXPEJoObedu4qOV5eUOnLwK0bAsEGlsfmh/ycSSN7yhPWOnXY+XuvLv9lqLt37Np4UjIlPRrRWuWc6t3yMOw6U9IBwR6kM1YUPzyr43NsyonBYi3myZEFR2Tg6kblH8RUURW5p8fNqnAQ9q9oIR89qZr/A/jtRcfEVyGK6nIhinsubCwzj1z0nuV828zKlkdbEwesyRhBgZD3g4QMNA6edToG9mewJorN2hJyWdeTya7fLsniFlSoEHc0lpUCS4FS7gFwmRit2bfFby4lxHc7AxSr/XTogY21l/ozI9A2KRuqEPetLwB9ZWvnXg94wjVXvPV3nlRfaJPZOIa51PjgWmd9VYOhVQ85nG1Px+WWkTREta4y1gpHfl8yRDWvHwu1/YSW28IBdcABrfbyNMpU2cz0LHSRhjJp0h24Kc4+k1OJdMH47fykl5qmQMSvzJnbx/+UP8+jCJDvMtgCcBJNo8CGmkDHQRHKYk7CHt64pW/84g1Q3gAIrdegTwj6mR9y5JbPt/F5/ySGdEAEPPeg7NHn9s2prIcr#3509697703`
	deText, err := DecryptFile(ver20230909, "test.lua", []byte(enText))
	if err != nil {
		t.Fatal(err)
	}
	oriText := `local Tools = luanet.import_type('JyGame.Tools')
local LuaTool = luanet.import_type('JyGame.LuaTool')
local CommonSettings = luanet.import_type('JyGame.CommonSettings')
local RuntimeData = luanet.import_type('JyGame.RuntimeData')
function AI_GetAIResult(battleLogic,rst)
return nil
end
function AI_GetEmptyBlockResult(battlefield, currentSprite, skill, x, y)
if (Tools.ProbabilityTest(0.5)) then
if (skill.Name == "召唤树妖" or skill.Name == "召唤金刚" or skill.Name == "召唤罗汉" or skill.Name == "召唤尸鬼") then
return Tools.GetRandom(1000, 10000)
end
if (skill.Name == "还魂咒") then
if (((currentSprite.Team == 1 and battlefield:GetTeam1DeadCount() > 0) or (currentSprite.Team == 2 and battlefield:GetTeam2DeadCount() > 0)) and Tools.ProbabilityTest(0.6)) then
return Tools.GetRandom(10000, math.max(currentSprite.MaxHp, 20000))
end
end
end
return 0
end
-- 由softmgr提供技术支持`
	if string(deText) != oriText {
		t.Fatal("解密失败")
	}
}

func TestVer20230909Xml(t *testing.T) {
	enText := `@KO8X9FjMbsDbVEZgDnv0nDshEtQMbg2xczI2+lL2nDs2/PH/Q7F+F70zi5e8Xe0zoDzAnIbZsuOGNuuDsUTIvDOymgjRiEyZks56FnO6awrJKwJEYLdtjCULSRqKrNBRtHziHFUMyFp7Y2sm9CR970Nx4RHu5g3s3mfIf18ESV3ty0L+ECoFVYkiNz1i0Ns/Qhubw66wyOSSZG0+jAGdOgEThO7RnyfEUSdvA4vx0I6FRzW7ejfJ24Qi4FPaRnjPgLGq1yG1VVUqFTeTSJy0ZMz3nI4JadRi+/MCt2lwGd1NQFGkcQ+Cu2SAm98fVU3XzWzI7IGyLxs7KG7jcSe/60wVgE14x5hFEbMYbmdhSKup6idVam1juIN1dwxuCV5kLTraKcN5NGAXQmzn9kBon8HTOF83sE+uBnTnME+T51WzjK8BB5Xu3IoRP4iwhBiCxBumnHcvH75U2oy5UK9Hb2NlAzwvraoBPvKHJt/aMSo=#3225521577`
	deText, err := DecryptFile(ver20230909, "test.xml", []byte(enText))
	if err != nil {
		t.Fatal(err)
	}
	oriText := `<root>
<map name="南贤居死胡同2" pic="地图.帮派大厅1" desc="一个死胡同">
<mapunit name="南贤" pic="" description="美丽的白发大姐姐南贤" x="-1" y="-1">
<event description="" value="mainStory_online" repeat="" type="story" image="" probability="100" />
</mapunit>
<musics />
</map>
</root>
<!-- 由softmgr提供技术支持 -->`
	if string(deText) != oriText {
		t.Fatal("解密失败")
	}
}
