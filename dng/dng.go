package dng

import "github.com/jonathanpittman/tiff"

var (
	DNGv1_0_0_0Tags = tiff.NewTagSet("DNGv1.0.0.0", 32768, 65535)
	DNGv1_1_0_0Tags = tiff.NewTagSet("DNGv1.1.0.0", 32768, 65535)
	DNGv1_2_0_0Tags = tiff.NewTagSet("DNGv1.2.0.0", 32768, 65535)
	DNGv1_3_0_0Tags = tiff.NewTagSet("DNGv1.3.0.0", 32768, 65535)
	DNGv1_4_0_0Tags = tiff.NewTagSet("DNGv1.4.0.0", 32768, 65535)
)

func init() {
	// Original tags for Version 1.0.0.0
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50706, "DNGVersion", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50707, "DNGBackwardVersion", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50708, "UniqueCameraModel", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50709, "LocalizedCameraModel", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50710, "CFAPlaneColor", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50711, "CFALayout", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50712, "LinearizationTable", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50713, "BlackLevelRepeatDim", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50714, "BlackLevel", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50715, "BlackLevelDeltaH", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50716, "BlackLevelDeltaV", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50717, "WhiteLevel", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50718, "DefaultScale", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50780, "BestQualityScale", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50719, "DefaultCropOrigin", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50720, "DefaultCropSize", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50778, "CalibrationIlluminant1", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50779, "CalibrationIlluminant2", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50721, "ColorMatrix1", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50722, "ColorMatrix2", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50723, "CameraCalibration1", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50724, "CameraCalibration2", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50725, "ReductionMatrix1", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50726, "ReductionMatrix2", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50727, "AnalogBalance", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50728, "AsShotNeutral", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50729, "AsShotWhiteXY", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50730, "BaselineExposure", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50731, "BaselineNoise", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50732, "BaselineSharpness", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50733, "BayerGreenSplit", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50734, "LinearResponseLimit", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50735, "CameraSerialNumber", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50736, "LensInfo", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50737, "ChromaBlurRadius", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50738, "AntiAliasStrength", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50740, "DNGPrivateData", nil))
	DNGv1_0_0_0Tags.Register(tiff.NewTag(50741, "MakerNoteSafety", nil))
	DNGv1_0_0_0Tags.Lock()
	tiff.DefaultTagSpace.RegisterTagSet(DNGv1_0_0_0Tags)

	// Additional Tags for Version 1.1.0.0
	DNGv1_1_0_0Tags.Register(tiff.NewTag(50739, "ShadowScale", nil))
	DNGv1_1_0_0Tags.Register(tiff.NewTag(50781, "RawDataUniqueID", nil))
	DNGv1_1_0_0Tags.Register(tiff.NewTag(50827, "OriginalRawFileName", nil))
	DNGv1_1_0_0Tags.Register(tiff.NewTag(50828, "OriginalRawFileData", nil))
	DNGv1_1_0_0Tags.Register(tiff.NewTag(50829, "ActiveArea", nil))
	DNGv1_1_0_0Tags.Register(tiff.NewTag(50830, "MaskedAreas", nil))
	DNGv1_1_0_0Tags.Register(tiff.NewTag(50831, "AsShotICCProfile", nil))
	DNGv1_1_0_0Tags.Register(tiff.NewTag(50832, "AsShotPreProfileMatrix", nil))
	DNGv1_1_0_0Tags.Register(tiff.NewTag(50833, "CurrentICCProfile", nil))
	DNGv1_1_0_0Tags.Register(tiff.NewTag(50834, "CurrentPreProfileMatrix", nil))
	DNGv1_1_0_0Tags.Lock()
	tiff.DefaultTagSpace.RegisterTagSet(DNGv1_1_0_0Tags)

	// Additional Tags for Version 1.2.0.0
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50879, "ColorimetricReference", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50931, "CameraCalibrationSignature", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50932, "ProfileCalibrationSignature", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50933, "ExtraCameraProfiles", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50934, "AsShotProfileName", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50935, "NoiseReductionApplied", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50936, "ProfileName", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50937, "ProfileHueSatMapDims", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50938, "ProfileHueSatMapData1", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50939, "ProfileHueSatMapData2", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50940, "ProfileToneCurve", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50941, "ProfileEmbedPolicy", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50942, "ProfileCopyright", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50964, "ForwardMatrix1", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50965, "ForwardMatrix2", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50966, "PreviewApplicationName", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50967, "PreviewApplicationVersion", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50968, "PreviewSettingsName", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50969, "PreviewSettingsDigest", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50970, "PreviewColorSpace", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50971, "PreviewDateTime", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50972, "RawImageDigest", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50973, "OriginalRawFileDigest", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50974, "SubTileBlockSize", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50975, "RowInterleaveFactor", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50981, "ProfileLookTableDims", nil))
	DNGv1_2_0_0Tags.Register(tiff.NewTag(50982, "ProfileLookTableData", nil))
	DNGv1_2_0_0Tags.Lock()
	tiff.DefaultTagSpace.RegisterTagSet(DNGv1_2_0_0Tags)

	// Additional Tags for Version 1.3.0.0
	DNGv1_3_0_0Tags.Register(tiff.NewTag(51008, "OpcodeList1", nil))
	DNGv1_3_0_0Tags.Register(tiff.NewTag(51009, "OpcodeList2", nil))
	DNGv1_3_0_0Tags.Register(tiff.NewTag(51022, "OpcodeList3", nil))
	DNGv1_3_0_0Tags.Register(tiff.NewTag(51041, "NoiseProfile", nil))
	DNGv1_3_0_0Tags.Lock()
	tiff.DefaultTagSpace.RegisterTagSet(DNGv1_3_0_0Tags)

	// Additional Tags for Version 1.4.0.0
	DNGv1_4_0_0Tags.Register(tiff.NewTag(51125, "DefaultUserCrop", nil))
	DNGv1_4_0_0Tags.Register(tiff.NewTag(51110, "DefaultBlackRender", nil))
	DNGv1_4_0_0Tags.Register(tiff.NewTag(51109, "BaselineExposureOffset", nil))
	DNGv1_4_0_0Tags.Register(tiff.NewTag(51108, "ProfileLookTableEncoding", nil))
	DNGv1_4_0_0Tags.Register(tiff.NewTag(51107, "ProfileHueSatMapEncoding", nil))
	DNGv1_4_0_0Tags.Register(tiff.NewTag(51089, "OriginalDefaultFinalSize", nil))
	DNGv1_4_0_0Tags.Register(tiff.NewTag(51090, "OriginalBestQualityFinalSize", nil))
	DNGv1_4_0_0Tags.Register(tiff.NewTag(51091, "OriginalDefaultCropSize", nil))
	DNGv1_4_0_0Tags.Register(tiff.NewTag(51111, "NewRawImageDigest", nil))
	DNGv1_4_0_0Tags.Register(tiff.NewTag(51112, "RawToPreviewGain", nil))
	DNGv1_4_0_0Tags.Lock()
	tiff.DefaultTagSpace.RegisterTagSet(DNGv1_4_0_0Tags)
}