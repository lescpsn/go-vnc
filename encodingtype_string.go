// Code generated by "stringer -type=EncodingType"; DO NOT EDIT.

package vnc

import "fmt"

const _EncodingType_name = "EncContinuousUpdatesPseudoEncFencePseudoEncClientRedirectEncXvpPseudoEncExtendedDesktopSizePseudoEncDesktopNamePseudoEncTightPngEncQEMUExtendedKeyEventPseudoEncQEMUPointerMotionChangePseudoEncCompressionLevel1EncCompressionLevel2EncCompressionLevel3EncCompressionLevel4EncCompressionLevel5EncCompressionLevel6EncCompressionLevel7EncCompressionLevel8EncCompressionLevel9EncCompressionLevel10EncXCursorPseudoEncCursorPseudoEncLastRectPseudoEncDesktopSizePseudoEncJPEGQualityLevelPseudo1EncJPEGQualityLevelPseudo2EncJPEGQualityLevelPseudo3EncJPEGQualityLevelPseudo4EncJPEGQualityLevelPseudo5EncJPEGQualityLevelPseudo6EncJPEGQualityLevelPseudo7EncJPEGQualityLevelPseudo8EncJPEGQualityLevelPseudo9EncJPEGQualityLevelPseudo10EncRawEncCopyRectEncRREEncCoRREEncHextileEncZlibEncTightEncZlibHexEncUltra1EncUltra2EncTRLEEncZRLEEncJPEGEncJRLEEncAtenAST2100EncAtenASTJPEGEncAtenHermonEncAtenYarkonEncAtenPilot3"

var _EncodingType_map = map[EncodingType]string{
	-313: _EncodingType_name[0:26],
	-312: _EncodingType_name[26:40],
	-311: _EncodingType_name[40:57],
	-309: _EncodingType_name[57:69],
	-308: _EncodingType_name[69:97],
	-307: _EncodingType_name[97:117],
	-260: _EncodingType_name[117:128],
	-258: _EncodingType_name[128:157],
	-257: _EncodingType_name[157:189],
	-256: _EncodingType_name[189:209],
	-255: _EncodingType_name[209:229],
	-254: _EncodingType_name[229:249],
	-253: _EncodingType_name[249:269],
	-252: _EncodingType_name[269:289],
	-251: _EncodingType_name[289:309],
	-250: _EncodingType_name[309:329],
	-249: _EncodingType_name[329:349],
	-248: _EncodingType_name[349:369],
	-247: _EncodingType_name[369:390],
	-240: _EncodingType_name[390:406],
	-239: _EncodingType_name[406:421],
	-224: _EncodingType_name[421:438],
	-223: _EncodingType_name[438:458],
	-32:  _EncodingType_name[458:484],
	-31:  _EncodingType_name[484:510],
	-30:  _EncodingType_name[510:536],
	-29:  _EncodingType_name[536:562],
	-28:  _EncodingType_name[562:588],
	-27:  _EncodingType_name[588:614],
	-26:  _EncodingType_name[614:640],
	-25:  _EncodingType_name[640:666],
	-24:  _EncodingType_name[666:692],
	-23:  _EncodingType_name[692:719],
	0:    _EncodingType_name[719:725],
	1:    _EncodingType_name[725:736],
	2:    _EncodingType_name[736:742],
	4:    _EncodingType_name[742:750],
	5:    _EncodingType_name[750:760],
	6:    _EncodingType_name[760:767],
	7:    _EncodingType_name[767:775],
	8:    _EncodingType_name[775:785],
	9:    _EncodingType_name[785:794],
	10:   _EncodingType_name[794:803],
	15:   _EncodingType_name[803:810],
	16:   _EncodingType_name[810:817],
	21:   _EncodingType_name[817:824],
	22:   _EncodingType_name[824:831],
	87:   _EncodingType_name[831:845],
	88:   _EncodingType_name[845:859],
	89:   _EncodingType_name[859:872],
	96:   _EncodingType_name[872:885],
	97:   _EncodingType_name[885:898],
}

func (i EncodingType) String() string {
	if str, ok := _EncodingType_map[i]; ok {
		return str
	}
	return fmt.Sprintf("EncodingType(%d)", i)
}
