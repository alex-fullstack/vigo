package matroska

const 	CRCIdValue = 0xbf

const 	EBMLMaxIDLengthDefault = 4

const (
	WrongType EMBLIDType = iota
	EBMLHeaderIdType
	EBMLVersionIdType
	EBMLReadVersionIdType
	EBMLMaxIDLengthIdType
	EBMLMaxSizeLengthIdType
	DocTypeIdType
	DocTypeVersionIdType
	DocTypeReadVersionIdType
	SegmentIdType
	EBMLEmptyIdType
	SeekHeadIdType
	SeekIdType
	SeekIDIdType
	SeekPositionIdType
	InfoIdType
	SegmentUIDIdType
	SegmentFilenameIdType
	PrevUIDIdType
	PrevFilenameIdType
	NextUIDIdType
	NextFilenameIdType
	SegmentFamilyIdType
	ChapterTranslateIdType
	ChapterTranslateEditionUIDIdType
	ChapterTranslateCodecIdType
	ChapterTranslateIDIdType
	TimecodeScaleIdType
	DurationIdType
	DateUTCIdType
	TitleIdType
	MuxingAppIdType
	WritingAppIdType
	ClusterIdType
	TimecodeIdType
	SilentTracksIdType
	SilentTrackNumberIdType
	PositionIdType
	PrevSizeIdType
	SimpleBlockIdType
	BlockGroupIdType
	BlockIdType
	BlockVirtualIdType
	BlockAdditionsIdType
	BlockMoreIdType
	BlockAddIDIdType
	BlockAdditionalIdType
	BlockDurationIdType
	ReferencePriorityIdType
	ReferenceBlockIdType
	ReferenceVirtualIdType
	CodecStateIdType
	DiscardPaddingIdType
	SlicesIdType
	TimeSliceIdType
	LaceNumberIdType
	FrameNumberIdType
	BlockAdditionIDIdType
	DelayIdType
	SliceDurationIdType
	ReferenceFrameIdType
	ReferenceOffsetIdType
	ReferenceTimeCodeIdType
	EncryptedBlockIdType
	TracksIdType
	TrackEntryIdType
	TrackNumberIdType
	TrackUIDIdType
	TrackTypeIdType
	FlagEnabledIdType
	FlagDefaultIdType
	FlagForcedIdType
	FlagLacingIdType
	MinCacheIdType
	MaxCacheIdType
	DefaultDurationIdType
	DefaultDecodedFieldDurationIdType
	TrackTimecodeScaleIdType
	TrackOffsetIdType
	MaxBlockAdditionIDIdType
	NameIdType
	LanguageIdType
	CodecIDIdType
	CodecPrivateIdType
	CodecNameIdType
	AttachmentLinkIdType
	CodecSettingsIdType
	CodecInfoURLIdType
	CodecDownloadURLIdType
	CodecDecodeAllIdType
	TrackOverlayIdType
	CodecDelayIdType
	SeekPreRollIdType
	TrackTranslateIdType
	TrackTranslateEditionUIDIdType
	TrackTranslateCodecIdType
	TrackTranslateTrackIDIdType
	VideoIdType
	FlagInterlacedIdType
	StereoModeIdType
	AlphaModeIdType
	OldStereoModeIdType
	PixelWidthIdType
	PixelHeightIdType
	PixelCropBottomIdType
	PixelCropTopIdType
	PixelCropLeftIdType
	PixelCropRightIdType
	DisplayWidthIdType
	DisplayHeightIdType
	DisplayUnitIdType
	AspectRatioTypeIdType
	ColourSpaceIdType
	GammaValueIdType
	FrameRateIdType
	AudioIdType
	SamplingFrequencyIdType
	OutputSamplingFrequencyIdType
	ChannelsIdType
	ChannelPositionsIdType
	BitDepthIdType
	TrackOperationIdType
	TrackCombinePlanesIdType
	TrackPlaneIdType
	TrackPlaneUIDIdType
	TrackPlaneTypeIdType
	TrackJoinBlocksIdType
	TrackJoinUIDIdType
	TrickTrackUIDIdType
	TrickTrackSegmentUIDIdType
	TrickTrackFlagIdType
	TrickMasterTrackUIDIdType
	TrickMasterTrackSegmentUIDIdType
	ContentEncodingsIdType
	ContentEncodingIdType
	ContentEncodingOrderIdType
	ContentEncodingScopeIdType
	ContentEncodingTypeIdType
	ContentCompressionIdType
	ContentCompAlgoIdType
	ContentCompSettingsIdType
	ContentEncryptionIdType
	ContentEncAlgoIdType
	ContentEncKeyIDIdType
	ContentSignatureIdType
	ContentSigKeyIDIdType
	ContentSigAlgoIdType
	ContentSigHashAlgoIdType
	CuesIdType
	CuePointIdType
	CueTimeIdType
	CueTrackPositionsIdType
	CueTrackIdType
	CueClusterPositionIdType
	CueRelativePositionIdType
	CueDurationIdType
	CueBlockNumberIdType
	CueCodecStateIdType
	CueReferenceIdType
	CueRefTimeIdType
	CueRefClusterIdType
	CueRefNumberIdType
	CueRefCodecStateIdType
	AttachmentsIdType
	AttachedFileIdType
	FileDescriptionIdType
	FileNameIdType
	FileMimeTypeIdType
	FileDataIdType
	FileUIDIdType
	FileReferralIdType
	FileUsedStartTimeIdType
	FileUsedEndTimeIdType
	ChaptersIdType
	EditionEntryIdType
	EditionUIDIdType
	EditionFlagHiddenIdType
	EditionFlagDefaultIdType
	EditionFlagOrderedIdType
	ChapterAtomIdType
	ChapterUIDIdType
	ChapterStringUIDIdType
	ChapterTimeStartIdType
	ChapterTimeEndIdType
	ChapterFlagHiddenIdType
	ChapterFlagEnabledIdType
	ChapterSegmentUIDIdType
	ChapterSegmentEditionUIDIdType
	ChapterPhysicalEquivIdType
	ChapterTrackIdType
	ChapterTrackNumberIdType
	ChapterDisplayIdType
	ChapStringIdType
	ChapLanguageIdType
	ChapCountryIdType
	ChapProcessIdType
	ChapProcessCodecIDIdType
	ChapProcessPrivateIdType
	ChapProcessCommandIdType
	ChapProcessTimeIdType
	ChapProcessDataIdType
	TagsIdType
	TagIdType
	TargetsIdType
	TargetTypeValueIdType
	TargetTypeIdType
	TagTrackUIDIdType
	TagEditionUIDIdType
	TagChapterUIDIdType
	TagAttachmentUIDIdType
	SimpleTagIdType
	TagNameIdType
	TagLanguageIdType
	TagDefaultIdType
	TagStringIdType
	TagBinaryIdType
)
const (
	EBMLHeaderId EMBLIDValue = 0x1A45DFA3
	EBMLVersionId EMBLIDValue = 0x4286
	EBMLReadVersionId EMBLIDValue = 0x42F7
	EBMLMaxIDLengthId EMBLIDValue = 0x42F2
	EBMLMaxSizeLengthId EMBLIDValue = 0x42F3
	DocTypeId EMBLIDValue = 0x4282
	DocTypeVersionId EMBLIDValue = 0x4287
	DocTypeReadVersionId EMBLIDValue = 0x4285
	SegmentId EMBLIDValue = 0x18538067
	EBMLEmptyId EMBLIDValue = 0xEC
	SeekHeadId EMBLIDValue = 0x114D9B74
	SeekId EMBLIDValue = 0x4DBB
	SeekIDId EMBLIDValue = 0x53AB
	SeekPositionId EMBLIDValue = 0x53AC
	InfoId EMBLIDValue = 0x1549A966
	SegmentUIDId EMBLIDValue = 0x73A4
	SegmentFilenameId EMBLIDValue = 0x7384
	PrevUIDId EMBLIDValue = 0x3CB923
	PrevFilenameId EMBLIDValue = 0x3C83AB
	NextUIDId EMBLIDValue = 0x3EB923
	NextFilenameId EMBLIDValue = 0x3E83BB
	SegmentFamilyId EMBLIDValue = 0x4444
	ChapterTranslateId EMBLIDValue = 0x6924
	ChapterTranslateEditionUIDId EMBLIDValue = 0x69FC
	ChapterTranslateCodecId EMBLIDValue = 0x69BF
	ChapterTranslateIDId EMBLIDValue = 0x69A5
	TimecodeScaleId EMBLIDValue = 0x2AD7B1
	DurationId EMBLIDValue = 0x4489
	DateUTCId EMBLIDValue = 0x4461
	TitleId EMBLIDValue = 0x7BA9
	MuxingAppId EMBLIDValue = 0x4D80
	WritingAppId EMBLIDValue = 0x5741
	ClusterId EMBLIDValue = 0x1F43B675
	TimecodeId EMBLIDValue = 0xE7
	SilentTracksId EMBLIDValue = 0x5854
	SilentTrackNumberId EMBLIDValue = 0x58D7
	PositionId EMBLIDValue = 0xA7
	PrevSizeId EMBLIDValue = 0xAB
	SimpleBlockId EMBLIDValue = 0xA3
	BlockGroupId EMBLIDValue = 0xA0
	BlockId EMBLIDValue = 0xA1
	BlockVirtualId EMBLIDValue = 0xA2
	BlockAdditionsId EMBLIDValue = 0x75A1
	BlockMoreId EMBLIDValue = 0xA6
	BlockAddIDId EMBLIDValue = 0xEE
	BlockAdditionalId EMBLIDValue = 0xA5
	BlockDurationId EMBLIDValue = 0x9B
	ReferencePriorityId EMBLIDValue = 0xFA
	ReferenceBlockId EMBLIDValue = 0xFB
	ReferenceVirtualId EMBLIDValue = 0xFD
	CodecStateId EMBLIDValue = 0xA4
	DiscardPaddingId EMBLIDValue = 0x75A2
	SlicesId EMBLIDValue = 0x8E
	TimeSliceId EMBLIDValue = 0xE8
	LaceNumberId EMBLIDValue = 0xCC
	FrameNumberId EMBLIDValue = 0xCD
	BlockAdditionIDId EMBLIDValue = 0xCB
	DelayId EMBLIDValue = 0xCE
	SliceDurationId EMBLIDValue = 0xCF
	ReferenceFrameId EMBLIDValue = 0xC8
	ReferenceOffsetId EMBLIDValue = 0xC9
	ReferenceTimeCodeId EMBLIDValue = 0xCA
	EncryptedBlockId EMBLIDValue = 0xAF
	TracksId EMBLIDValue = 0x1654AE6B
	TrackEntryId EMBLIDValue = 0xAE
	TrackNumberId EMBLIDValue = 0xD7
	TrackUIDId EMBLIDValue = 0x73C5
	TrackTypeId EMBLIDValue = 0x83
	FlagEnabledId EMBLIDValue = 0xB9
	FlagDefaultId EMBLIDValue = 0x88
	FlagForcedId EMBLIDValue = 0x55AA
	FlagLacingId EMBLIDValue = 0x9C
	MinCacheId EMBLIDValue = 0x6DE7
	MaxCacheId EMBLIDValue = 0x6DF8
	DefaultDurationId EMBLIDValue = 0x23E383
	DefaultDecodedFieldDurationId EMBLIDValue = 0x234E7A
	TrackTimecodeScaleId EMBLIDValue = 0x23314F
	TrackOffsetId EMBLIDValue = 0x537F
	MaxBlockAdditionIDId EMBLIDValue = 0x55EE
	NameId EMBLIDValue = 0x536E
	LanguageId EMBLIDValue = 0x22B59C
	CodecIDId EMBLIDValue = 0x86
	CodecPrivateId EMBLIDValue = 0x63A2
	CodecNameId EMBLIDValue = 0x258688
	AttachmentLinkId EMBLIDValue = 0x7446
	CodecSettingsId EMBLIDValue = 0x3A9697
	CodecInfoURLId EMBLIDValue = 0x3B4040
	CodecDownloadURLId EMBLIDValue = 0x26B240
	CodecDecodeAllId EMBLIDValue = 0xAA
	TrackOverlayId EMBLIDValue = 0x6FAB
	CodecDelayId EMBLIDValue = 0x56AA
	SeekPreRollId EMBLIDValue = 0x56BB
	TrackTranslateId EMBLIDValue = 0x6624
	TrackTranslateEditionUIDId EMBLIDValue = 0x66FC
	TrackTranslateCodecId EMBLIDValue = 0x66BF
	TrackTranslateTrackIDId EMBLIDValue = 0x66A5
	VideoId EMBLIDValue = 0xE0
	FlagInterlacedId EMBLIDValue = 0x9A
	StereoModeId EMBLIDValue = 0x53B8
	AlphaModeId EMBLIDValue = 0x53C0
	OldStereoModeId EMBLIDValue = 0x53B9
	PixelWidthId EMBLIDValue = 0xB0
	PixelHeightId EMBLIDValue = 0xBA
	PixelCropBottomId EMBLIDValue = 0x54AA
	PixelCropTopId EMBLIDValue = 0x54BB
	PixelCropLeftId EMBLIDValue = 0x54CC
	PixelCropRightId EMBLIDValue = 0x54DD
	DisplayWidthId EMBLIDValue = 0x54B0
	DisplayHeightId EMBLIDValue = 0x54BA
	DisplayUnitId EMBLIDValue = 0x54B2
	AspectRatioTypeId EMBLIDValue = 0x54B3
	ColourSpaceId EMBLIDValue = 0x2EB524
	GammaValueId EMBLIDValue = 0x2FB523
	FrameRateId EMBLIDValue = 0x2383E3
	AudioId EMBLIDValue = 0xE1
	SamplingFrequencyId EMBLIDValue = 0xB5
	OutputSamplingFrequencyId EMBLIDValue = 0x78B5
	ChannelsId EMBLIDValue = 0x9F
	ChannelPositionsId EMBLIDValue = 0x7D7B
	BitDepthId EMBLIDValue = 0x6264
	TrackOperationId EMBLIDValue = 0xE2
	TrackCombinePlanesId EMBLIDValue = 0xE3
	TrackPlaneId EMBLIDValue = 0xE4
	TrackPlaneUIDId EMBLIDValue = 0xE5
	TrackPlaneTypeId EMBLIDValue = 0xE6
	TrackJoinBlocksId EMBLIDValue = 0xE9
	TrackJoinUIDId EMBLIDValue = 0xED
	TrickTrackUIDId EMBLIDValue = 0xC0
	TrickTrackSegmentUIDId EMBLIDValue = 0xC1
	TrickTrackFlagId EMBLIDValue = 0xC6
	TrickMasterTrackUIDId EMBLIDValue = 0xC7
	TrickMasterTrackSegmentUIDId EMBLIDValue = 0xC4
	ContentEncodingsId EMBLIDValue = 0x6D80
	ContentEncodingId EMBLIDValue = 0x6240
	ContentEncodingOrderId EMBLIDValue = 0x5031
	ContentEncodingScopeId EMBLIDValue = 0x5032
	ContentEncodingTypeId EMBLIDValue = 0x5033
	ContentCompressionId EMBLIDValue = 0x5034
	ContentCompAlgoId EMBLIDValue = 0x4254
	ContentCompSettingsId EMBLIDValue = 0x4255
	ContentEncryptionId EMBLIDValue = 0x5035
	ContentEncAlgoId EMBLIDValue = 0x47E1
	ContentEncKeyIDId EMBLIDValue = 0x47E2
	ContentSignatureId EMBLIDValue = 0x47E3
	ContentSigKeyIDId EMBLIDValue = 0x47E4
	ContentSigAlgoId EMBLIDValue = 0x47E5
	ContentSigHashAlgoId EMBLIDValue = 0x47E6
	CuesId EMBLIDValue = 0x1C53BB6B
	CuePointId EMBLIDValue = 0xBB
	CueTimeId EMBLIDValue = 0xB3
	CueTrackPositionsId EMBLIDValue = 0xB7
	CueTrackId EMBLIDValue = 0xF7
	CueClusterPositionId EMBLIDValue = 0xF1
	CueRelativePositionId EMBLIDValue = 0xF0
	CueDurationId EMBLIDValue = 0xB2
	CueBlockNumberId EMBLIDValue = 0x5378
	CueCodecStateId EMBLIDValue = 0xEA
	CueReferenceId EMBLIDValue = 0xDB
	CueRefTimeId EMBLIDValue = 0x96
	CueRefClusterId EMBLIDValue = 0x97
	CueRefNumberId EMBLIDValue = 0x535F
	CueRefCodecStateId EMBLIDValue = 0xEB
	AttachmentsId EMBLIDValue = 0x1941A469
	AttachedFileId EMBLIDValue = 0x61A7
	FileDescriptionId EMBLIDValue = 0x467E
	FileNameId EMBLIDValue = 0x466E
	FileMimeTypeId EMBLIDValue = 0x4660
	FileDataId EMBLIDValue = 0x465C
	FileUIDId EMBLIDValue = 0x46AE
	FileReferralId EMBLIDValue = 0x4675
	FileUsedStartTimeId EMBLIDValue = 0x4661
	FileUsedEndTimeId EMBLIDValue = 0x4662
	ChaptersId EMBLIDValue = 0x1043A770
	EditionEntryId EMBLIDValue = 0x45B9
	EditionUIDId EMBLIDValue = 0x45BC
	EditionFlagHiddenId EMBLIDValue = 0x45BD
	EditionFlagDefaultId EMBLIDValue = 0x45DB
	EditionFlagOrderedId EMBLIDValue = 0x45DD
	ChapterAtomId EMBLIDValue = 0xB6
	ChapterUIDId EMBLIDValue = 0x73C4
	ChapterStringUIDId EMBLIDValue = 0x5654
	ChapterTimeStartId EMBLIDValue = 0x91
	ChapterTimeEndId EMBLIDValue = 0x92
	ChapterFlagHiddenId EMBLIDValue = 0x98
	ChapterFlagEnabledId EMBLIDValue = 0x4598
	ChapterSegmentUIDId EMBLIDValue = 0x6E67
	ChapterSegmentEditionUIDId EMBLIDValue = 0x6EBC
	ChapterPhysicalEquivId EMBLIDValue = 0x63C3
	ChapterTrackId EMBLIDValue = 0x8F
	ChapterTrackNumberId EMBLIDValue = 0x89
	ChapterDisplayId EMBLIDValue = 0x80
	ChapStringId EMBLIDValue = 0x85
	ChapLanguageId EMBLIDValue = 0x437C
	ChapCountryId EMBLIDValue = 0x437E
	ChapProcessId EMBLIDValue = 0x6944
	ChapProcessCodecIDId EMBLIDValue = 0x6955
	ChapProcessPrivateId EMBLIDValue = 0x450D
	ChapProcessCommandId EMBLIDValue = 0x6911
	ChapProcessTimeId EMBLIDValue = 0x6922
	ChapProcessDataId EMBLIDValue = 0x6933
	TagsId EMBLIDValue = 0x1254C367
	TagId EMBLIDValue = 0x7373
	TargetsId EMBLIDValue = 0x63C0
	TargetTypeValueId EMBLIDValue = 0x68CA
	TargetTypeId EMBLIDValue = 0x63CA
	TagTrackUIDId EMBLIDValue = 0x63C5
	TagEditionUIDId EMBLIDValue = 0x63C9
	TagChapterUIDId EMBLIDValue = 0x63C4
	TagAttachmentUIDId EMBLIDValue = 0x63C6
	SimpleTagId EMBLIDValue = 0x67C8
	TagNameId EMBLIDValue = 0x45A3
	TagLanguageId EMBLIDValue = 0x447A
	TagDefaultId EMBLIDValue = 0x4484
	TagStringId EMBLIDValue = 0x4487
	TagBinaryId EMBLIDValue = 0x4485
)

type EMBLIDType uint

type EMBLIDValue int

type SpecificationEBMLHeader struct {
	SpecificationElement  
	EBMLVersion  *SpecificationElement  `json:"eBMLVersion,omitempty"`
	EBMLReadVersion  *SpecificationElement  `json:"eBMLReadVersion,omitempty"`
	EBMLMaxIDLength  *SpecificationElement  `json:"eBMLMaxIDLength,omitempty"`
	EBMLMaxSizeLength  *SpecificationElement  `json:"eBMLMaxSizeLength,omitempty"`
	DocType  *SpecificationElement  `json:"docType,omitempty"`
	DocTypeVersion  *SpecificationElement  `json:"docTypeVersion,omitempty"`
	DocTypeReadVersion  *SpecificationElement  `json:"docTypeReadVersion,omitempty"`
}
func newSpecificationEBMLHeader() *SpecificationEBMLHeader {
	return &SpecificationEBMLHeader{}
}

type SpecificationSeek struct {
	SpecificationElement  
	SeekID  *SpecificationElement  `json:"seekID,omitempty"`
	SeekPosition  *SpecificationElement  `json:"seekPosition,omitempty"`
}
func newSpecificationSeek() *SpecificationSeek {
	return &SpecificationSeek{}
}

type SpecificationSeekHead struct {
	SpecificationElement  
	Seek  []*SpecificationSeek  `json:"seek,omitempty"`
}
func newSpecificationSeekHead() *SpecificationSeekHead {
	return &SpecificationSeekHead{}
}

type SpecificationChapterTranslate struct {
	SpecificationElement  
	ChapterTranslateEditionUID  []*SpecificationElement  `json:"chapterTranslateEditionUID,omitempty"`
	ChapterTranslateCodec  *SpecificationElement  `json:"chapterTranslateCodec,omitempty"`
	ChapterTranslateID  *SpecificationElement  `json:"chapterTranslateID,omitempty"`
}
func newSpecificationChapterTranslate() *SpecificationChapterTranslate {
	return &SpecificationChapterTranslate{}
}

type SpecificationInfo struct {
	SpecificationElement  
	SegmentUID  *SpecificationElement  `json:"segmentUID,omitempty"`
	SegmentFilename  *SpecificationElement  `json:"segmentFilename,omitempty"`
	PrevUID  *SpecificationElement  `json:"prevUID,omitempty"`
	PrevFilename  *SpecificationElement  `json:"prevFilename,omitempty"`
	NextUID  *SpecificationElement  `json:"nextUID,omitempty"`
	NextFilename  *SpecificationElement  `json:"nextFilename,omitempty"`
	SegmentFamily  []*SpecificationElement  `json:"segmentFamily,omitempty"`
	ChapterTranslate  []*SpecificationChapterTranslate  `json:"chapterTranslate,omitempty"`
	TimecodeScale  *SpecificationElement  `json:"timecodeScale,omitempty"`
	Duration  *SpecificationElement  `json:"duration,omitempty"`
	DateUTC  *SpecificationElement  `json:"dateUTC,omitempty"`
	Title  *SpecificationElement  `json:"title,omitempty"`
	MuxingApp  *SpecificationElement  `json:"muxingApp,omitempty"`
	WritingApp  *SpecificationElement  `json:"writingApp,omitempty"`
}
func newSpecificationInfo() *SpecificationInfo {
	return &SpecificationInfo{}
}

type SpecificationSilentTracks struct {
	SpecificationElement  
	SilentTrackNumber  []*SpecificationElement  `json:"silentTrackNumber,omitempty"`
}
func newSpecificationSilentTracks() *SpecificationSilentTracks {
	return &SpecificationSilentTracks{}
}

type SpecificationBlockMore struct {
	SpecificationElement  
	BlockAddID  *SpecificationElement  `json:"blockAddID,omitempty"`
	BlockAdditional  *SpecificationElement  `json:"blockAdditional,omitempty"`
}
func newSpecificationBlockMore() *SpecificationBlockMore {
	return &SpecificationBlockMore{}
}

type SpecificationBlockAdditions struct {
	SpecificationElement  
	BlockMore  []*SpecificationBlockMore  `json:"blockMore,omitempty"`
}
func newSpecificationBlockAdditions() *SpecificationBlockAdditions {
	return &SpecificationBlockAdditions{}
}

type SpecificationTimeSlice struct {
	SpecificationElement  
	LaceNumber  *SpecificationElement  `json:"laceNumber,omitempty"`
	FrameNumber  *SpecificationElement  `json:"frameNumber,omitempty"`
	BlockAdditionID  *SpecificationElement  `json:"blockAdditionID,omitempty"`
	Delay  *SpecificationElement  `json:"delay,omitempty"`
	SliceDuration  *SpecificationElement  `json:"sliceDuration,omitempty"`
}
func newSpecificationTimeSlice() *SpecificationTimeSlice {
	return &SpecificationTimeSlice{}
}

type SpecificationSlices struct {
	SpecificationElement  
	TimeSlice  []*SpecificationTimeSlice  `json:"timeSlice,omitempty"`
}
func newSpecificationSlices() *SpecificationSlices {
	return &SpecificationSlices{}
}

type SpecificationReferenceFrame struct {
	SpecificationElement  
	ReferenceOffset  *SpecificationElement  `json:"referenceOffset,omitempty"`
	ReferenceTimeCode  *SpecificationElement  `json:"referenceTimeCode,omitempty"`
}
func newSpecificationReferenceFrame() *SpecificationReferenceFrame {
	return &SpecificationReferenceFrame{}
}

type SpecificationBlockGroup struct {
	SpecificationElement  
	Block  *SpecificationElement  `json:"block,omitempty"`
	BlockVirtual  *SpecificationElement  `json:"blockVirtual,omitempty"`
	BlockAdditions  *SpecificationBlockAdditions  `json:"blockAdditions,omitempty"`
	BlockDuration  *SpecificationElement  `json:"blockDuration,omitempty"`
	ReferencePriority  *SpecificationElement  `json:"referencePriority,omitempty"`
	ReferenceBlock  []*SpecificationElement  `json:"referenceBlock,omitempty"`
	ReferenceVirtual  *SpecificationElement  `json:"referenceVirtual,omitempty"`
	CodecState  *SpecificationElement  `json:"codecState,omitempty"`
	DiscardPadding  *SpecificationElement  `json:"discardPadding,omitempty"`
	Slices  *SpecificationSlices  `json:"slices,omitempty"`
	ReferenceFrame  *SpecificationReferenceFrame  `json:"referenceFrame,omitempty"`
}
func newSpecificationBlockGroup() *SpecificationBlockGroup {
	return &SpecificationBlockGroup{}
}

type SpecificationCluster struct {
	SpecificationElement  
	Timecode  *SpecificationElement  `json:"timecode,omitempty"`
	SilentTracks  *SpecificationSilentTracks  `json:"silentTracks,omitempty"`
	Position  *SpecificationElement  `json:"position,omitempty"`
	PrevSize  *SpecificationElement  `json:"prevSize,omitempty"`
	SimpleBlock  []*SpecificationElement  `json:"simpleBlock,omitempty"`
	BlockGroup  []*SpecificationBlockGroup  `json:"blockGroup,omitempty"`
	EncryptedBlock  []*SpecificationElement  `json:"encryptedBlock,omitempty"`
}
func newSpecificationCluster() *SpecificationCluster {
	return &SpecificationCluster{}
}

type SpecificationTrackTranslate struct {
	SpecificationElement  
	TrackTranslateEditionUID  []*SpecificationElement  `json:"trackTranslateEditionUID,omitempty"`
	TrackTranslateCodec  *SpecificationElement  `json:"trackTranslateCodec,omitempty"`
	TrackTranslateTrackID  *SpecificationElement  `json:"trackTranslateTrackID,omitempty"`
}
func newSpecificationTrackTranslate() *SpecificationTrackTranslate {
	return &SpecificationTrackTranslate{}
}

type SpecificationVideo struct {
	SpecificationElement  
	FlagInterlaced  *SpecificationElement  `json:"flagInterlaced,omitempty"`
	StereoMode  *SpecificationElement  `json:"stereoMode,omitempty"`
	AlphaMode  *SpecificationElement  `json:"alphaMode,omitempty"`
	OldStereoMode  *SpecificationElement  `json:"oldStereoMode,omitempty"`
	PixelWidth  *SpecificationElement  `json:"pixelWidth,omitempty"`
	PixelHeight  *SpecificationElement  `json:"pixelHeight,omitempty"`
	PixelCropBottom  *SpecificationElement  `json:"pixelCropBottom,omitempty"`
	PixelCropTop  *SpecificationElement  `json:"pixelCropTop,omitempty"`
	PixelCropLeft  *SpecificationElement  `json:"pixelCropLeft,omitempty"`
	PixelCropRight  *SpecificationElement  `json:"pixelCropRight,omitempty"`
	DisplayWidth  *SpecificationElement  `json:"displayWidth,omitempty"`
	DisplayHeight  *SpecificationElement  `json:"displayHeight,omitempty"`
	DisplayUnit  *SpecificationElement  `json:"displayUnit,omitempty"`
	AspectRatioType  *SpecificationElement  `json:"aspectRatioType,omitempty"`
	ColourSpace  *SpecificationElement  `json:"colourSpace,omitempty"`
	GammaValue  *SpecificationElement  `json:"gammaValue,omitempty"`
	FrameRate  *SpecificationElement  `json:"frameRate,omitempty"`
}
func newSpecificationVideo() *SpecificationVideo {
	return &SpecificationVideo{}
}

type SpecificationAudio struct {
	SpecificationElement  
	SamplingFrequency  *SpecificationElement  `json:"samplingFrequency,omitempty"`
	OutputSamplingFrequency  *SpecificationElement  `json:"outputSamplingFrequency,omitempty"`
	Channels  *SpecificationElement  `json:"channels,omitempty"`
	ChannelPositions  *SpecificationElement  `json:"channelPositions,omitempty"`
	BitDepth  *SpecificationElement  `json:"bitDepth,omitempty"`
}
func newSpecificationAudio() *SpecificationAudio {
	return &SpecificationAudio{}
}

type SpecificationTrackPlane struct {
	SpecificationElement  
	TrackPlaneUID  *SpecificationElement  `json:"trackPlaneUID,omitempty"`
	TrackPlaneType  *SpecificationElement  `json:"trackPlaneType,omitempty"`
}
func newSpecificationTrackPlane() *SpecificationTrackPlane {
	return &SpecificationTrackPlane{}
}

type SpecificationTrackCombinePlanes struct {
	SpecificationElement  
	TrackPlane  []*SpecificationTrackPlane  `json:"trackPlane,omitempty"`
}
func newSpecificationTrackCombinePlanes() *SpecificationTrackCombinePlanes {
	return &SpecificationTrackCombinePlanes{}
}

type SpecificationTrackJoinBlocks struct {
	SpecificationElement  
	TrackJoinUID  []*SpecificationElement  `json:"trackJoinUID,omitempty"`
}
func newSpecificationTrackJoinBlocks() *SpecificationTrackJoinBlocks {
	return &SpecificationTrackJoinBlocks{}
}

type SpecificationTrackOperation struct {
	SpecificationElement  
	TrackCombinePlanes  *SpecificationTrackCombinePlanes  `json:"trackCombinePlanes,omitempty"`
	TrackJoinBlocks  *SpecificationTrackJoinBlocks  `json:"trackJoinBlocks,omitempty"`
}
func newSpecificationTrackOperation() *SpecificationTrackOperation {
	return &SpecificationTrackOperation{}
}

type SpecificationContentCompression struct {
	SpecificationElement  
	ContentCompAlgo  *SpecificationElement  `json:"contentCompAlgo,omitempty"`
	ContentCompSettings  *SpecificationElement  `json:"contentCompSettings,omitempty"`
}
func newSpecificationContentCompression() *SpecificationContentCompression {
	return &SpecificationContentCompression{}
}

type SpecificationContentEncryption struct {
	SpecificationElement  
	ContentEncAlgo  *SpecificationElement  `json:"contentEncAlgo,omitempty"`
	ContentEncKeyID  *SpecificationElement  `json:"contentEncKeyID,omitempty"`
	ContentSignature  *SpecificationElement  `json:"contentSignature,omitempty"`
	ContentSigKeyID  *SpecificationElement  `json:"contentSigKeyID,omitempty"`
	ContentSigAlgo  *SpecificationElement  `json:"contentSigAlgo,omitempty"`
	ContentSigHashAlgo  *SpecificationElement  `json:"contentSigHashAlgo,omitempty"`
}
func newSpecificationContentEncryption() *SpecificationContentEncryption {
	return &SpecificationContentEncryption{}
}

type SpecificationContentEncoding struct {
	SpecificationElement  
	ContentEncodingOrder  *SpecificationElement  `json:"contentEncodingOrder,omitempty"`
	ContentEncodingScope  *SpecificationElement  `json:"contentEncodingScope,omitempty"`
	ContentEncodingType  *SpecificationElement  `json:"contentEncodingType,omitempty"`
	ContentCompression  *SpecificationContentCompression  `json:"contentCompression,omitempty"`
	ContentEncryption  *SpecificationContentEncryption  `json:"contentEncryption,omitempty"`
}
func newSpecificationContentEncoding() *SpecificationContentEncoding {
	return &SpecificationContentEncoding{}
}

type SpecificationContentEncodings struct {
	SpecificationElement  
	ContentEncoding  []*SpecificationContentEncoding  `json:"contentEncoding,omitempty"`
}
func newSpecificationContentEncodings() *SpecificationContentEncodings {
	return &SpecificationContentEncodings{}
}

type SpecificationTrackEntry struct {
	SpecificationElement  
	TrackNumber  *SpecificationElement  `json:"trackNumber,omitempty"`
	TrackUID  *SpecificationElement  `json:"trackUID,omitempty"`
	TrackType  *SpecificationElement  `json:"trackType,omitempty"`
	FlagEnabled  *SpecificationElement  `json:"flagEnabled,omitempty"`
	FlagDefault  *SpecificationElement  `json:"flagDefault,omitempty"`
	FlagForced  *SpecificationElement  `json:"flagForced,omitempty"`
	FlagLacing  *SpecificationElement  `json:"flagLacing,omitempty"`
	MinCache  *SpecificationElement  `json:"minCache,omitempty"`
	MaxCache  *SpecificationElement  `json:"maxCache,omitempty"`
	DefaultDuration  *SpecificationElement  `json:"defaultDuration,omitempty"`
	DefaultDecodedFieldDuration  *SpecificationElement  `json:"defaultDecodedFieldDuration,omitempty"`
	TrackTimecodeScale  *SpecificationElement  `json:"trackTimecodeScale,omitempty"`
	TrackOffset  *SpecificationElement  `json:"trackOffset,omitempty"`
	MaxBlockAdditionID  *SpecificationElement  `json:"maxBlockAdditionID,omitempty"`
	Name  *SpecificationElement  `json:"name,omitempty"`
	Language  *SpecificationElement  `json:"language,omitempty"`
	CodecID  *SpecificationElement  `json:"codecID,omitempty"`
	CodecPrivate  *SpecificationElement  `json:"codecPrivate,omitempty"`
	CodecName  *SpecificationElement  `json:"codecName,omitempty"`
	AttachmentLink  *SpecificationElement  `json:"attachmentLink,omitempty"`
	CodecSettings  *SpecificationElement  `json:"codecSettings,omitempty"`
	CodecInfoURL  []*SpecificationElement  `json:"codecInfoURL,omitempty"`
	CodecDownloadURL  []*SpecificationElement  `json:"codecDownloadURL,omitempty"`
	CodecDecodeAll  *SpecificationElement  `json:"codecDecodeAll,omitempty"`
	TrackOverlay  []*SpecificationElement  `json:"trackOverlay,omitempty"`
	CodecDelay  *SpecificationElement  `json:"codecDelay,omitempty"`
	SeekPreRoll  *SpecificationElement  `json:"seekPreRoll,omitempty"`
	TrackTranslate  []*SpecificationTrackTranslate  `json:"trackTranslate,omitempty"`
	Video  *SpecificationVideo  `json:"video,omitempty"`
	Audio  *SpecificationAudio  `json:"audio,omitempty"`
	TrackOperation  *SpecificationTrackOperation  `json:"trackOperation,omitempty"`
	TrickTrackUID  *SpecificationElement  `json:"trickTrackUID,omitempty"`
	TrickTrackSegmentUID  *SpecificationElement  `json:"trickTrackSegmentUID,omitempty"`
	TrickTrackFlag  *SpecificationElement  `json:"trickTrackFlag,omitempty"`
	TrickMasterTrackUID  *SpecificationElement  `json:"trickMasterTrackUID,omitempty"`
	TrickMasterTrackSegmentUID  *SpecificationElement  `json:"trickMasterTrackSegmentUID,omitempty"`
	ContentEncodings  *SpecificationContentEncodings  `json:"contentEncodings,omitempty"`
}
func newSpecificationTrackEntry() *SpecificationTrackEntry {
	return &SpecificationTrackEntry{}
}

type SpecificationTracks struct {
	SpecificationElement  
	TrackEntry  []*SpecificationTrackEntry  `json:"trackEntry,omitempty"`
}
func newSpecificationTracks() *SpecificationTracks {
	return &SpecificationTracks{}
}

type SpecificationCueReference struct {
	SpecificationElement  
	CueRefTime  *SpecificationElement  `json:"cueRefTime,omitempty"`
	CueRefCluster  *SpecificationElement  `json:"cueRefCluster,omitempty"`
	CueRefNumber  *SpecificationElement  `json:"cueRefNumber,omitempty"`
	CueRefCodecState  *SpecificationElement  `json:"cueRefCodecState,omitempty"`
}
func newSpecificationCueReference() *SpecificationCueReference {
	return &SpecificationCueReference{}
}

type SpecificationCueTrackPositions struct {
	SpecificationElement  
	CueTrack  *SpecificationElement  `json:"cueTrack,omitempty"`
	CueClusterPosition  *SpecificationElement  `json:"cueClusterPosition,omitempty"`
	CueRelativePosition  *SpecificationElement  `json:"cueRelativePosition,omitempty"`
	CueDuration  *SpecificationElement  `json:"cueDuration,omitempty"`
	CueBlockNumber  *SpecificationElement  `json:"cueBlockNumber,omitempty"`
	CueCodecState  *SpecificationElement  `json:"cueCodecState,omitempty"`
	CueReference  []*SpecificationCueReference  `json:"cueReference,omitempty"`
}
func newSpecificationCueTrackPositions() *SpecificationCueTrackPositions {
	return &SpecificationCueTrackPositions{}
}

type SpecificationCuePoint struct {
	SpecificationElement  
	CueTime  *SpecificationElement  `json:"cueTime,omitempty"`
	CueTrackPositions  []*SpecificationCueTrackPositions  `json:"cueTrackPositions,omitempty"`
}
func newSpecificationCuePoint() *SpecificationCuePoint {
	return &SpecificationCuePoint{}
}

type SpecificationCues struct {
	SpecificationElement  
	CuePoint  []*SpecificationCuePoint  `json:"cuePoint,omitempty"`
}
func newSpecificationCues() *SpecificationCues {
	return &SpecificationCues{}
}

type SpecificationAttachedFile struct {
	SpecificationElement  
	FileDescription  *SpecificationElement  `json:"fileDescription,omitempty"`
	FileName  *SpecificationElement  `json:"fileName,omitempty"`
	FileMimeType  *SpecificationElement  `json:"fileMimeType,omitempty"`
	FileData  *SpecificationElement  `json:"fileData,omitempty"`
	FileUID  *SpecificationElement  `json:"fileUID,omitempty"`
	FileReferral  *SpecificationElement  `json:"fileReferral,omitempty"`
	FileUsedStartTime  *SpecificationElement  `json:"fileUsedStartTime,omitempty"`
	FileUsedEndTime  *SpecificationElement  `json:"fileUsedEndTime,omitempty"`
}
func newSpecificationAttachedFile() *SpecificationAttachedFile {
	return &SpecificationAttachedFile{}
}

type SpecificationAttachments struct {
	SpecificationElement  
	AttachedFile  []*SpecificationAttachedFile  `json:"attachedFile,omitempty"`
}
func newSpecificationAttachments() *SpecificationAttachments {
	return &SpecificationAttachments{}
}

type SpecificationChapterTrack struct {
	SpecificationElement  
	ChapterTrackNumber  []*SpecificationElement  `json:"chapterTrackNumber,omitempty"`
}
func newSpecificationChapterTrack() *SpecificationChapterTrack {
	return &SpecificationChapterTrack{}
}

type SpecificationChapterDisplay struct {
	SpecificationElement  
	ChapString  *SpecificationElement  `json:"chapString,omitempty"`
	ChapLanguage  []*SpecificationElement  `json:"chapLanguage,omitempty"`
	ChapCountry  []*SpecificationElement  `json:"chapCountry,omitempty"`
}
func newSpecificationChapterDisplay() *SpecificationChapterDisplay {
	return &SpecificationChapterDisplay{}
}

type SpecificationChapProcessCommand struct {
	SpecificationElement  
	ChapProcessTime  *SpecificationElement  `json:"chapProcessTime,omitempty"`
	ChapProcessData  *SpecificationElement  `json:"chapProcessData,omitempty"`
}
func newSpecificationChapProcessCommand() *SpecificationChapProcessCommand {
	return &SpecificationChapProcessCommand{}
}

type SpecificationChapProcess struct {
	SpecificationElement  
	ChapProcessCodecID  *SpecificationElement  `json:"chapProcessCodecID,omitempty"`
	ChapProcessPrivate  *SpecificationElement  `json:"chapProcessPrivate,omitempty"`
	ChapProcessCommand  []*SpecificationChapProcessCommand  `json:"chapProcessCommand,omitempty"`
}
func newSpecificationChapProcess() *SpecificationChapProcess {
	return &SpecificationChapProcess{}
}

type SpecificationChapterAtom struct {
	SpecificationElement  
	ChapterUID  *SpecificationElement  `json:"chapterUID,omitempty"`
	ChapterStringUID  *SpecificationElement  `json:"chapterStringUID,omitempty"`
	ChapterTimeStart  *SpecificationElement  `json:"chapterTimeStart,omitempty"`
	ChapterTimeEnd  *SpecificationElement  `json:"chapterTimeEnd,omitempty"`
	ChapterFlagHidden  *SpecificationElement  `json:"chapterFlagHidden,omitempty"`
	ChapterFlagEnabled  *SpecificationElement  `json:"chapterFlagEnabled,omitempty"`
	ChapterSegmentUID  *SpecificationElement  `json:"chapterSegmentUID,omitempty"`
	ChapterSegmentEditionUID  *SpecificationElement  `json:"chapterSegmentEditionUID,omitempty"`
	ChapterPhysicalEquiv  *SpecificationElement  `json:"chapterPhysicalEquiv,omitempty"`
	ChapterTrack  *SpecificationChapterTrack  `json:"chapterTrack,omitempty"`
	ChapterDisplay  []*SpecificationChapterDisplay  `json:"chapterDisplay,omitempty"`
	ChapProcess  []*SpecificationChapProcess  `json:"chapProcess,omitempty"`
}
func newSpecificationChapterAtom() *SpecificationChapterAtom {
	return &SpecificationChapterAtom{}
}

type SpecificationEditionEntry struct {
	SpecificationElement  
	EditionUID  *SpecificationElement  `json:"editionUID,omitempty"`
	EditionFlagHidden  *SpecificationElement  `json:"editionFlagHidden,omitempty"`
	EditionFlagDefault  *SpecificationElement  `json:"editionFlagDefault,omitempty"`
	EditionFlagOrdered  *SpecificationElement  `json:"editionFlagOrdered,omitempty"`
	ChapterAtom  []*SpecificationChapterAtom  `json:"chapterAtom,omitempty"`
}
func newSpecificationEditionEntry() *SpecificationEditionEntry {
	return &SpecificationEditionEntry{}
}

type SpecificationChapters struct {
	SpecificationElement  
	EditionEntry  []*SpecificationEditionEntry  `json:"editionEntry,omitempty"`
}
func newSpecificationChapters() *SpecificationChapters {
	return &SpecificationChapters{}
}

type SpecificationTargets struct {
	SpecificationElement  
	TargetTypeValue  *SpecificationElement  `json:"targetTypeValue,omitempty"`
	TargetType  *SpecificationElement  `json:"targetType,omitempty"`
	TagTrackUID  []*SpecificationElement  `json:"tagTrackUID,omitempty"`
	TagEditionUID  []*SpecificationElement  `json:"tagEditionUID,omitempty"`
	TagChapterUID  []*SpecificationElement  `json:"tagChapterUID,omitempty"`
	TagAttachmentUID  []*SpecificationElement  `json:"tagAttachmentUID,omitempty"`
}
func newSpecificationTargets() *SpecificationTargets {
	return &SpecificationTargets{}
}

type SpecificationSimpleTag struct {
	SpecificationElement  
	TagName  *SpecificationElement  `json:"tagName,omitempty"`
	TagLanguage  *SpecificationElement  `json:"tagLanguage,omitempty"`
	TagDefault  *SpecificationElement  `json:"tagDefault,omitempty"`
	TagString  *SpecificationElement  `json:"tagString,omitempty"`
	TagBinary  *SpecificationElement  `json:"tagBinary,omitempty"`
}
func newSpecificationSimpleTag() *SpecificationSimpleTag {
	return &SpecificationSimpleTag{}
}

type SpecificationTag struct {
	SpecificationElement  
	Targets  *SpecificationTargets  `json:"targets,omitempty"`
	SimpleTag  []*SpecificationSimpleTag  `json:"simpleTag,omitempty"`
}
func newSpecificationTag() *SpecificationTag {
	return &SpecificationTag{}
}

type SpecificationTags struct {
	SpecificationElement  
	Tag  []*SpecificationTag  `json:"tag,omitempty"`
}
func newSpecificationTags() *SpecificationTags {
	return &SpecificationTags{}
}

type SpecificationSegment struct {
	SpecificationElement  
	EBMLEmpty  *SpecificationElement  `json:"eBMLEmpty,omitempty"`
	SeekHead  []*SpecificationSeekHead  `json:"seekHead,omitempty"`
	Info  []*SpecificationInfo  `json:"info,omitempty"`
	Cluster  []*SpecificationCluster  `json:"cluster,omitempty"`
	Tracks  []*SpecificationTracks  `json:"tracks,omitempty"`
	Cues  *SpecificationCues  `json:"cues,omitempty"`
	Attachments  *SpecificationAttachments  `json:"attachments,omitempty"`
	Chapters  *SpecificationChapters  `json:"chapters,omitempty"`
	Tags  []*SpecificationTags  `json:"tags,omitempty"`
}
func newSpecificationSegment() *SpecificationSegment {
	return &SpecificationSegment{}
}

type Specification struct {
	EBMLHeader  []*SpecificationEBMLHeader  `json:"eBMLHeader,omitempty"`
	Segment  []*SpecificationSegment  `json:"segment,omitempty"`
}
func newSpecification() *Specification {
	return &Specification{}
}


func convertIdValueToIdType(value EMBLIDValue) EMBLIDType {
	switch value {
		case EBMLHeaderId:
			return EBMLHeaderIdType
		case EBMLVersionId:
			return EBMLVersionIdType
		case EBMLReadVersionId:
			return EBMLReadVersionIdType
		case EBMLMaxIDLengthId:
			return EBMLMaxIDLengthIdType
		case EBMLMaxSizeLengthId:
			return EBMLMaxSizeLengthIdType
		case DocTypeId:
			return DocTypeIdType
		case DocTypeVersionId:
			return DocTypeVersionIdType
		case DocTypeReadVersionId:
			return DocTypeReadVersionIdType
		case SegmentId:
			return SegmentIdType
		case EBMLEmptyId:
			return EBMLEmptyIdType
		case SeekHeadId:
			return SeekHeadIdType
		case SeekId:
			return SeekIdType
		case SeekIDId:
			return SeekIDIdType
		case SeekPositionId:
			return SeekPositionIdType
		case InfoId:
			return InfoIdType
		case SegmentUIDId:
			return SegmentUIDIdType
		case SegmentFilenameId:
			return SegmentFilenameIdType
		case PrevUIDId:
			return PrevUIDIdType
		case PrevFilenameId:
			return PrevFilenameIdType
		case NextUIDId:
			return NextUIDIdType
		case NextFilenameId:
			return NextFilenameIdType
		case SegmentFamilyId:
			return SegmentFamilyIdType
		case ChapterTranslateId:
			return ChapterTranslateIdType
		case ChapterTranslateEditionUIDId:
			return ChapterTranslateEditionUIDIdType
		case ChapterTranslateCodecId:
			return ChapterTranslateCodecIdType
		case ChapterTranslateIDId:
			return ChapterTranslateIDIdType
		case TimecodeScaleId:
			return TimecodeScaleIdType
		case DurationId:
			return DurationIdType
		case DateUTCId:
			return DateUTCIdType
		case TitleId:
			return TitleIdType
		case MuxingAppId:
			return MuxingAppIdType
		case WritingAppId:
			return WritingAppIdType
		case ClusterId:
			return ClusterIdType
		case TimecodeId:
			return TimecodeIdType
		case SilentTracksId:
			return SilentTracksIdType
		case SilentTrackNumberId:
			return SilentTrackNumberIdType
		case PositionId:
			return PositionIdType
		case PrevSizeId:
			return PrevSizeIdType
		case SimpleBlockId:
			return SimpleBlockIdType
		case BlockGroupId:
			return BlockGroupIdType
		case BlockId:
			return BlockIdType
		case BlockVirtualId:
			return BlockVirtualIdType
		case BlockAdditionsId:
			return BlockAdditionsIdType
		case BlockMoreId:
			return BlockMoreIdType
		case BlockAddIDId:
			return BlockAddIDIdType
		case BlockAdditionalId:
			return BlockAdditionalIdType
		case BlockDurationId:
			return BlockDurationIdType
		case ReferencePriorityId:
			return ReferencePriorityIdType
		case ReferenceBlockId:
			return ReferenceBlockIdType
		case ReferenceVirtualId:
			return ReferenceVirtualIdType
		case CodecStateId:
			return CodecStateIdType
		case DiscardPaddingId:
			return DiscardPaddingIdType
		case SlicesId:
			return SlicesIdType
		case TimeSliceId:
			return TimeSliceIdType
		case LaceNumberId:
			return LaceNumberIdType
		case FrameNumberId:
			return FrameNumberIdType
		case BlockAdditionIDId:
			return BlockAdditionIDIdType
		case DelayId:
			return DelayIdType
		case SliceDurationId:
			return SliceDurationIdType
		case ReferenceFrameId:
			return ReferenceFrameIdType
		case ReferenceOffsetId:
			return ReferenceOffsetIdType
		case ReferenceTimeCodeId:
			return ReferenceTimeCodeIdType
		case EncryptedBlockId:
			return EncryptedBlockIdType
		case TracksId:
			return TracksIdType
		case TrackEntryId:
			return TrackEntryIdType
		case TrackNumberId:
			return TrackNumberIdType
		case TrackUIDId:
			return TrackUIDIdType
		case TrackTypeId:
			return TrackTypeIdType
		case FlagEnabledId:
			return FlagEnabledIdType
		case FlagDefaultId:
			return FlagDefaultIdType
		case FlagForcedId:
			return FlagForcedIdType
		case FlagLacingId:
			return FlagLacingIdType
		case MinCacheId:
			return MinCacheIdType
		case MaxCacheId:
			return MaxCacheIdType
		case DefaultDurationId:
			return DefaultDurationIdType
		case DefaultDecodedFieldDurationId:
			return DefaultDecodedFieldDurationIdType
		case TrackTimecodeScaleId:
			return TrackTimecodeScaleIdType
		case TrackOffsetId:
			return TrackOffsetIdType
		case MaxBlockAdditionIDId:
			return MaxBlockAdditionIDIdType
		case NameId:
			return NameIdType
		case LanguageId:
			return LanguageIdType
		case CodecIDId:
			return CodecIDIdType
		case CodecPrivateId:
			return CodecPrivateIdType
		case CodecNameId:
			return CodecNameIdType
		case AttachmentLinkId:
			return AttachmentLinkIdType
		case CodecSettingsId:
			return CodecSettingsIdType
		case CodecInfoURLId:
			return CodecInfoURLIdType
		case CodecDownloadURLId:
			return CodecDownloadURLIdType
		case CodecDecodeAllId:
			return CodecDecodeAllIdType
		case TrackOverlayId:
			return TrackOverlayIdType
		case CodecDelayId:
			return CodecDelayIdType
		case SeekPreRollId:
			return SeekPreRollIdType
		case TrackTranslateId:
			return TrackTranslateIdType
		case TrackTranslateEditionUIDId:
			return TrackTranslateEditionUIDIdType
		case TrackTranslateCodecId:
			return TrackTranslateCodecIdType
		case TrackTranslateTrackIDId:
			return TrackTranslateTrackIDIdType
		case VideoId:
			return VideoIdType
		case FlagInterlacedId:
			return FlagInterlacedIdType
		case StereoModeId:
			return StereoModeIdType
		case AlphaModeId:
			return AlphaModeIdType
		case OldStereoModeId:
			return OldStereoModeIdType
		case PixelWidthId:
			return PixelWidthIdType
		case PixelHeightId:
			return PixelHeightIdType
		case PixelCropBottomId:
			return PixelCropBottomIdType
		case PixelCropTopId:
			return PixelCropTopIdType
		case PixelCropLeftId:
			return PixelCropLeftIdType
		case PixelCropRightId:
			return PixelCropRightIdType
		case DisplayWidthId:
			return DisplayWidthIdType
		case DisplayHeightId:
			return DisplayHeightIdType
		case DisplayUnitId:
			return DisplayUnitIdType
		case AspectRatioTypeId:
			return AspectRatioTypeIdType
		case ColourSpaceId:
			return ColourSpaceIdType
		case GammaValueId:
			return GammaValueIdType
		case FrameRateId:
			return FrameRateIdType
		case AudioId:
			return AudioIdType
		case SamplingFrequencyId:
			return SamplingFrequencyIdType
		case OutputSamplingFrequencyId:
			return OutputSamplingFrequencyIdType
		case ChannelsId:
			return ChannelsIdType
		case ChannelPositionsId:
			return ChannelPositionsIdType
		case BitDepthId:
			return BitDepthIdType
		case TrackOperationId:
			return TrackOperationIdType
		case TrackCombinePlanesId:
			return TrackCombinePlanesIdType
		case TrackPlaneId:
			return TrackPlaneIdType
		case TrackPlaneUIDId:
			return TrackPlaneUIDIdType
		case TrackPlaneTypeId:
			return TrackPlaneTypeIdType
		case TrackJoinBlocksId:
			return TrackJoinBlocksIdType
		case TrackJoinUIDId:
			return TrackJoinUIDIdType
		case TrickTrackUIDId:
			return TrickTrackUIDIdType
		case TrickTrackSegmentUIDId:
			return TrickTrackSegmentUIDIdType
		case TrickTrackFlagId:
			return TrickTrackFlagIdType
		case TrickMasterTrackUIDId:
			return TrickMasterTrackUIDIdType
		case TrickMasterTrackSegmentUIDId:
			return TrickMasterTrackSegmentUIDIdType
		case ContentEncodingsId:
			return ContentEncodingsIdType
		case ContentEncodingId:
			return ContentEncodingIdType
		case ContentEncodingOrderId:
			return ContentEncodingOrderIdType
		case ContentEncodingScopeId:
			return ContentEncodingScopeIdType
		case ContentEncodingTypeId:
			return ContentEncodingTypeIdType
		case ContentCompressionId:
			return ContentCompressionIdType
		case ContentCompAlgoId:
			return ContentCompAlgoIdType
		case ContentCompSettingsId:
			return ContentCompSettingsIdType
		case ContentEncryptionId:
			return ContentEncryptionIdType
		case ContentEncAlgoId:
			return ContentEncAlgoIdType
		case ContentEncKeyIDId:
			return ContentEncKeyIDIdType
		case ContentSignatureId:
			return ContentSignatureIdType
		case ContentSigKeyIDId:
			return ContentSigKeyIDIdType
		case ContentSigAlgoId:
			return ContentSigAlgoIdType
		case ContentSigHashAlgoId:
			return ContentSigHashAlgoIdType
		case CuesId:
			return CuesIdType
		case CuePointId:
			return CuePointIdType
		case CueTimeId:
			return CueTimeIdType
		case CueTrackPositionsId:
			return CueTrackPositionsIdType
		case CueTrackId:
			return CueTrackIdType
		case CueClusterPositionId:
			return CueClusterPositionIdType
		case CueRelativePositionId:
			return CueRelativePositionIdType
		case CueDurationId:
			return CueDurationIdType
		case CueBlockNumberId:
			return CueBlockNumberIdType
		case CueCodecStateId:
			return CueCodecStateIdType
		case CueReferenceId:
			return CueReferenceIdType
		case CueRefTimeId:
			return CueRefTimeIdType
		case CueRefClusterId:
			return CueRefClusterIdType
		case CueRefNumberId:
			return CueRefNumberIdType
		case CueRefCodecStateId:
			return CueRefCodecStateIdType
		case AttachmentsId:
			return AttachmentsIdType
		case AttachedFileId:
			return AttachedFileIdType
		case FileDescriptionId:
			return FileDescriptionIdType
		case FileNameId:
			return FileNameIdType
		case FileMimeTypeId:
			return FileMimeTypeIdType
		case FileDataId:
			return FileDataIdType
		case FileUIDId:
			return FileUIDIdType
		case FileReferralId:
			return FileReferralIdType
		case FileUsedStartTimeId:
			return FileUsedStartTimeIdType
		case FileUsedEndTimeId:
			return FileUsedEndTimeIdType
		case ChaptersId:
			return ChaptersIdType
		case EditionEntryId:
			return EditionEntryIdType
		case EditionUIDId:
			return EditionUIDIdType
		case EditionFlagHiddenId:
			return EditionFlagHiddenIdType
		case EditionFlagDefaultId:
			return EditionFlagDefaultIdType
		case EditionFlagOrderedId:
			return EditionFlagOrderedIdType
		case ChapterAtomId:
			return ChapterAtomIdType
		case ChapterUIDId:
			return ChapterUIDIdType
		case ChapterStringUIDId:
			return ChapterStringUIDIdType
		case ChapterTimeStartId:
			return ChapterTimeStartIdType
		case ChapterTimeEndId:
			return ChapterTimeEndIdType
		case ChapterFlagHiddenId:
			return ChapterFlagHiddenIdType
		case ChapterFlagEnabledId:
			return ChapterFlagEnabledIdType
		case ChapterSegmentUIDId:
			return ChapterSegmentUIDIdType
		case ChapterSegmentEditionUIDId:
			return ChapterSegmentEditionUIDIdType
		case ChapterPhysicalEquivId:
			return ChapterPhysicalEquivIdType
		case ChapterTrackId:
			return ChapterTrackIdType
		case ChapterTrackNumberId:
			return ChapterTrackNumberIdType
		case ChapterDisplayId:
			return ChapterDisplayIdType
		case ChapStringId:
			return ChapStringIdType
		case ChapLanguageId:
			return ChapLanguageIdType
		case ChapCountryId:
			return ChapCountryIdType
		case ChapProcessId:
			return ChapProcessIdType
		case ChapProcessCodecIDId:
			return ChapProcessCodecIDIdType
		case ChapProcessPrivateId:
			return ChapProcessPrivateIdType
		case ChapProcessCommandId:
			return ChapProcessCommandIdType
		case ChapProcessTimeId:
			return ChapProcessTimeIdType
		case ChapProcessDataId:
			return ChapProcessDataIdType
		case TagsId:
			return TagsIdType
		case TagId:
			return TagIdType
		case TargetsId:
			return TargetsIdType
		case TargetTypeValueId:
			return TargetTypeValueIdType
		case TargetTypeId:
			return TargetTypeIdType
		case TagTrackUIDId:
			return TagTrackUIDIdType
		case TagEditionUIDId:
			return TagEditionUIDIdType
		case TagChapterUIDId:
			return TagChapterUIDIdType
		case TagAttachmentUIDId:
			return TagAttachmentUIDIdType
		case SimpleTagId:
			return SimpleTagIdType
		case TagNameId:
			return TagNameIdType
		case TagLanguageId:
			return TagLanguageIdType
		case TagDefaultId:
			return TagDefaultIdType
		case TagStringId:
			return TagStringIdType
		case TagBinaryId:
			return TagBinaryIdType
		default:
			return WrongType
	}
}

func convertIdTypeToIdValue(value EMBLIDType) EMBLIDValue {
	switch value {
		case EBMLHeaderIdType:
			return EBMLHeaderId
		case EBMLVersionIdType:
			return EBMLVersionId
		case EBMLReadVersionIdType:
			return EBMLReadVersionId
		case EBMLMaxIDLengthIdType:
			return EBMLMaxIDLengthId
		case EBMLMaxSizeLengthIdType:
			return EBMLMaxSizeLengthId
		case DocTypeIdType:
			return DocTypeId
		case DocTypeVersionIdType:
			return DocTypeVersionId
		case DocTypeReadVersionIdType:
			return DocTypeReadVersionId
		case SegmentIdType:
			return SegmentId
		case EBMLEmptyIdType:
			return EBMLEmptyId
		case SeekHeadIdType:
			return SeekHeadId
		case SeekIdType:
			return SeekId
		case SeekIDIdType:
			return SeekIDId
		case SeekPositionIdType:
			return SeekPositionId
		case InfoIdType:
			return InfoId
		case SegmentUIDIdType:
			return SegmentUIDId
		case SegmentFilenameIdType:
			return SegmentFilenameId
		case PrevUIDIdType:
			return PrevUIDId
		case PrevFilenameIdType:
			return PrevFilenameId
		case NextUIDIdType:
			return NextUIDId
		case NextFilenameIdType:
			return NextFilenameId
		case SegmentFamilyIdType:
			return SegmentFamilyId
		case ChapterTranslateIdType:
			return ChapterTranslateId
		case ChapterTranslateEditionUIDIdType:
			return ChapterTranslateEditionUIDId
		case ChapterTranslateCodecIdType:
			return ChapterTranslateCodecId
		case ChapterTranslateIDIdType:
			return ChapterTranslateIDId
		case TimecodeScaleIdType:
			return TimecodeScaleId
		case DurationIdType:
			return DurationId
		case DateUTCIdType:
			return DateUTCId
		case TitleIdType:
			return TitleId
		case MuxingAppIdType:
			return MuxingAppId
		case WritingAppIdType:
			return WritingAppId
		case ClusterIdType:
			return ClusterId
		case TimecodeIdType:
			return TimecodeId
		case SilentTracksIdType:
			return SilentTracksId
		case SilentTrackNumberIdType:
			return SilentTrackNumberId
		case PositionIdType:
			return PositionId
		case PrevSizeIdType:
			return PrevSizeId
		case SimpleBlockIdType:
			return SimpleBlockId
		case BlockGroupIdType:
			return BlockGroupId
		case BlockIdType:
			return BlockId
		case BlockVirtualIdType:
			return BlockVirtualId
		case BlockAdditionsIdType:
			return BlockAdditionsId
		case BlockMoreIdType:
			return BlockMoreId
		case BlockAddIDIdType:
			return BlockAddIDId
		case BlockAdditionalIdType:
			return BlockAdditionalId
		case BlockDurationIdType:
			return BlockDurationId
		case ReferencePriorityIdType:
			return ReferencePriorityId
		case ReferenceBlockIdType:
			return ReferenceBlockId
		case ReferenceVirtualIdType:
			return ReferenceVirtualId
		case CodecStateIdType:
			return CodecStateId
		case DiscardPaddingIdType:
			return DiscardPaddingId
		case SlicesIdType:
			return SlicesId
		case TimeSliceIdType:
			return TimeSliceId
		case LaceNumberIdType:
			return LaceNumberId
		case FrameNumberIdType:
			return FrameNumberId
		case BlockAdditionIDIdType:
			return BlockAdditionIDId
		case DelayIdType:
			return DelayId
		case SliceDurationIdType:
			return SliceDurationId
		case ReferenceFrameIdType:
			return ReferenceFrameId
		case ReferenceOffsetIdType:
			return ReferenceOffsetId
		case ReferenceTimeCodeIdType:
			return ReferenceTimeCodeId
		case EncryptedBlockIdType:
			return EncryptedBlockId
		case TracksIdType:
			return TracksId
		case TrackEntryIdType:
			return TrackEntryId
		case TrackNumberIdType:
			return TrackNumberId
		case TrackUIDIdType:
			return TrackUIDId
		case TrackTypeIdType:
			return TrackTypeId
		case FlagEnabledIdType:
			return FlagEnabledId
		case FlagDefaultIdType:
			return FlagDefaultId
		case FlagForcedIdType:
			return FlagForcedId
		case FlagLacingIdType:
			return FlagLacingId
		case MinCacheIdType:
			return MinCacheId
		case MaxCacheIdType:
			return MaxCacheId
		case DefaultDurationIdType:
			return DefaultDurationId
		case DefaultDecodedFieldDurationIdType:
			return DefaultDecodedFieldDurationId
		case TrackTimecodeScaleIdType:
			return TrackTimecodeScaleId
		case TrackOffsetIdType:
			return TrackOffsetId
		case MaxBlockAdditionIDIdType:
			return MaxBlockAdditionIDId
		case NameIdType:
			return NameId
		case LanguageIdType:
			return LanguageId
		case CodecIDIdType:
			return CodecIDId
		case CodecPrivateIdType:
			return CodecPrivateId
		case CodecNameIdType:
			return CodecNameId
		case AttachmentLinkIdType:
			return AttachmentLinkId
		case CodecSettingsIdType:
			return CodecSettingsId
		case CodecInfoURLIdType:
			return CodecInfoURLId
		case CodecDownloadURLIdType:
			return CodecDownloadURLId
		case CodecDecodeAllIdType:
			return CodecDecodeAllId
		case TrackOverlayIdType:
			return TrackOverlayId
		case CodecDelayIdType:
			return CodecDelayId
		case SeekPreRollIdType:
			return SeekPreRollId
		case TrackTranslateIdType:
			return TrackTranslateId
		case TrackTranslateEditionUIDIdType:
			return TrackTranslateEditionUIDId
		case TrackTranslateCodecIdType:
			return TrackTranslateCodecId
		case TrackTranslateTrackIDIdType:
			return TrackTranslateTrackIDId
		case VideoIdType:
			return VideoId
		case FlagInterlacedIdType:
			return FlagInterlacedId
		case StereoModeIdType:
			return StereoModeId
		case AlphaModeIdType:
			return AlphaModeId
		case OldStereoModeIdType:
			return OldStereoModeId
		case PixelWidthIdType:
			return PixelWidthId
		case PixelHeightIdType:
			return PixelHeightId
		case PixelCropBottomIdType:
			return PixelCropBottomId
		case PixelCropTopIdType:
			return PixelCropTopId
		case PixelCropLeftIdType:
			return PixelCropLeftId
		case PixelCropRightIdType:
			return PixelCropRightId
		case DisplayWidthIdType:
			return DisplayWidthId
		case DisplayHeightIdType:
			return DisplayHeightId
		case DisplayUnitIdType:
			return DisplayUnitId
		case AspectRatioTypeIdType:
			return AspectRatioTypeId
		case ColourSpaceIdType:
			return ColourSpaceId
		case GammaValueIdType:
			return GammaValueId
		case FrameRateIdType:
			return FrameRateId
		case AudioIdType:
			return AudioId
		case SamplingFrequencyIdType:
			return SamplingFrequencyId
		case OutputSamplingFrequencyIdType:
			return OutputSamplingFrequencyId
		case ChannelsIdType:
			return ChannelsId
		case ChannelPositionsIdType:
			return ChannelPositionsId
		case BitDepthIdType:
			return BitDepthId
		case TrackOperationIdType:
			return TrackOperationId
		case TrackCombinePlanesIdType:
			return TrackCombinePlanesId
		case TrackPlaneIdType:
			return TrackPlaneId
		case TrackPlaneUIDIdType:
			return TrackPlaneUIDId
		case TrackPlaneTypeIdType:
			return TrackPlaneTypeId
		case TrackJoinBlocksIdType:
			return TrackJoinBlocksId
		case TrackJoinUIDIdType:
			return TrackJoinUIDId
		case TrickTrackUIDIdType:
			return TrickTrackUIDId
		case TrickTrackSegmentUIDIdType:
			return TrickTrackSegmentUIDId
		case TrickTrackFlagIdType:
			return TrickTrackFlagId
		case TrickMasterTrackUIDIdType:
			return TrickMasterTrackUIDId
		case TrickMasterTrackSegmentUIDIdType:
			return TrickMasterTrackSegmentUIDId
		case ContentEncodingsIdType:
			return ContentEncodingsId
		case ContentEncodingIdType:
			return ContentEncodingId
		case ContentEncodingOrderIdType:
			return ContentEncodingOrderId
		case ContentEncodingScopeIdType:
			return ContentEncodingScopeId
		case ContentEncodingTypeIdType:
			return ContentEncodingTypeId
		case ContentCompressionIdType:
			return ContentCompressionId
		case ContentCompAlgoIdType:
			return ContentCompAlgoId
		case ContentCompSettingsIdType:
			return ContentCompSettingsId
		case ContentEncryptionIdType:
			return ContentEncryptionId
		case ContentEncAlgoIdType:
			return ContentEncAlgoId
		case ContentEncKeyIDIdType:
			return ContentEncKeyIDId
		case ContentSignatureIdType:
			return ContentSignatureId
		case ContentSigKeyIDIdType:
			return ContentSigKeyIDId
		case ContentSigAlgoIdType:
			return ContentSigAlgoId
		case ContentSigHashAlgoIdType:
			return ContentSigHashAlgoId
		case CuesIdType:
			return CuesId
		case CuePointIdType:
			return CuePointId
		case CueTimeIdType:
			return CueTimeId
		case CueTrackPositionsIdType:
			return CueTrackPositionsId
		case CueTrackIdType:
			return CueTrackId
		case CueClusterPositionIdType:
			return CueClusterPositionId
		case CueRelativePositionIdType:
			return CueRelativePositionId
		case CueDurationIdType:
			return CueDurationId
		case CueBlockNumberIdType:
			return CueBlockNumberId
		case CueCodecStateIdType:
			return CueCodecStateId
		case CueReferenceIdType:
			return CueReferenceId
		case CueRefTimeIdType:
			return CueRefTimeId
		case CueRefClusterIdType:
			return CueRefClusterId
		case CueRefNumberIdType:
			return CueRefNumberId
		case CueRefCodecStateIdType:
			return CueRefCodecStateId
		case AttachmentsIdType:
			return AttachmentsId
		case AttachedFileIdType:
			return AttachedFileId
		case FileDescriptionIdType:
			return FileDescriptionId
		case FileNameIdType:
			return FileNameId
		case FileMimeTypeIdType:
			return FileMimeTypeId
		case FileDataIdType:
			return FileDataId
		case FileUIDIdType:
			return FileUIDId
		case FileReferralIdType:
			return FileReferralId
		case FileUsedStartTimeIdType:
			return FileUsedStartTimeId
		case FileUsedEndTimeIdType:
			return FileUsedEndTimeId
		case ChaptersIdType:
			return ChaptersId
		case EditionEntryIdType:
			return EditionEntryId
		case EditionUIDIdType:
			return EditionUIDId
		case EditionFlagHiddenIdType:
			return EditionFlagHiddenId
		case EditionFlagDefaultIdType:
			return EditionFlagDefaultId
		case EditionFlagOrderedIdType:
			return EditionFlagOrderedId
		case ChapterAtomIdType:
			return ChapterAtomId
		case ChapterUIDIdType:
			return ChapterUIDId
		case ChapterStringUIDIdType:
			return ChapterStringUIDId
		case ChapterTimeStartIdType:
			return ChapterTimeStartId
		case ChapterTimeEndIdType:
			return ChapterTimeEndId
		case ChapterFlagHiddenIdType:
			return ChapterFlagHiddenId
		case ChapterFlagEnabledIdType:
			return ChapterFlagEnabledId
		case ChapterSegmentUIDIdType:
			return ChapterSegmentUIDId
		case ChapterSegmentEditionUIDIdType:
			return ChapterSegmentEditionUIDId
		case ChapterPhysicalEquivIdType:
			return ChapterPhysicalEquivId
		case ChapterTrackIdType:
			return ChapterTrackId
		case ChapterTrackNumberIdType:
			return ChapterTrackNumberId
		case ChapterDisplayIdType:
			return ChapterDisplayId
		case ChapStringIdType:
			return ChapStringId
		case ChapLanguageIdType:
			return ChapLanguageId
		case ChapCountryIdType:
			return ChapCountryId
		case ChapProcessIdType:
			return ChapProcessId
		case ChapProcessCodecIDIdType:
			return ChapProcessCodecIDId
		case ChapProcessPrivateIdType:
			return ChapProcessPrivateId
		case ChapProcessCommandIdType:
			return ChapProcessCommandId
		case ChapProcessTimeIdType:
			return ChapProcessTimeId
		case ChapProcessDataIdType:
			return ChapProcessDataId
		case TagsIdType:
			return TagsId
		case TagIdType:
			return TagId
		case TargetsIdType:
			return TargetsId
		case TargetTypeValueIdType:
			return TargetTypeValueId
		case TargetTypeIdType:
			return TargetTypeId
		case TagTrackUIDIdType:
			return TagTrackUIDId
		case TagEditionUIDIdType:
			return TagEditionUIDId
		case TagChapterUIDIdType:
			return TagChapterUIDId
		case TagAttachmentUIDIdType:
			return TagAttachmentUIDId
		case SimpleTagIdType:
			return SimpleTagId
		case TagNameIdType:
			return TagNameId
		case TagLanguageIdType:
			return TagLanguageId
		case TagDefaultIdType:
			return TagDefaultId
		case TagStringIdType:
			return TagStringId
		case TagBinaryIdType:
			return TagBinaryId
		default:
			return -1
	}
}

func getElement(specification *Specification, value EMBLIDType) (el *SpecificationElement, hasChildren bool) {
	switch value {
		case EBMLHeaderIdType:
			if specification.EBMLHeader == nil {
				specification.EBMLHeader = []*SpecificationEBMLHeader{newSpecificationEBMLHeader()}
			} else {
				specification.EBMLHeader = append(specification.EBMLHeader, newSpecificationEBMLHeader())
			}
			return &specification.EBMLHeader[len(specification.EBMLHeader)-1].SpecificationElement, true
		case EBMLVersionIdType:
			if specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLVersion == nil {
				specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLVersion = newSpecificationElement()
			}
			return specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLVersion, false
		case EBMLReadVersionIdType:
			if specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLReadVersion == nil {
				specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLReadVersion = newSpecificationElement()
			}
			return specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLReadVersion, false
		case EBMLMaxIDLengthIdType:
			if specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLMaxIDLength == nil {
				specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLMaxIDLength = newSpecificationElement()
			}
			return specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLMaxIDLength, false
		case EBMLMaxSizeLengthIdType:
			if specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLMaxSizeLength == nil {
				specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLMaxSizeLength = newSpecificationElement()
			}
			return specification.EBMLHeader[len(specification.EBMLHeader)-1].EBMLMaxSizeLength, false
		case DocTypeIdType:
			if specification.EBMLHeader[len(specification.EBMLHeader)-1].DocType == nil {
				specification.EBMLHeader[len(specification.EBMLHeader)-1].DocType = newSpecificationElement()
			}
			return specification.EBMLHeader[len(specification.EBMLHeader)-1].DocType, false
		case DocTypeVersionIdType:
			if specification.EBMLHeader[len(specification.EBMLHeader)-1].DocTypeVersion == nil {
				specification.EBMLHeader[len(specification.EBMLHeader)-1].DocTypeVersion = newSpecificationElement()
			}
			return specification.EBMLHeader[len(specification.EBMLHeader)-1].DocTypeVersion, false
		case DocTypeReadVersionIdType:
			if specification.EBMLHeader[len(specification.EBMLHeader)-1].DocTypeReadVersion == nil {
				specification.EBMLHeader[len(specification.EBMLHeader)-1].DocTypeReadVersion = newSpecificationElement()
			}
			return specification.EBMLHeader[len(specification.EBMLHeader)-1].DocTypeReadVersion, false
		case SegmentIdType:
			if specification.Segment == nil {
				specification.Segment = []*SpecificationSegment{newSpecificationSegment()}
			} else {
				specification.Segment = append(specification.Segment, newSpecificationSegment())
			}
			return &specification.Segment[len(specification.Segment)-1].SpecificationElement, true
		case EBMLEmptyIdType:
			if specification.Segment[len(specification.Segment)-1].EBMLEmpty == nil {
				specification.Segment[len(specification.Segment)-1].EBMLEmpty = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].EBMLEmpty, false
		case SeekHeadIdType:
			if specification.Segment[len(specification.Segment)-1].SeekHead == nil {
				specification.Segment[len(specification.Segment)-1].SeekHead = []*SpecificationSeekHead{newSpecificationSeekHead()}
			} else {
				specification.Segment[len(specification.Segment)-1].SeekHead = append(specification.Segment[len(specification.Segment)-1].SeekHead, newSpecificationSeekHead())
			}
			return &specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].SpecificationElement, true
		case SeekIdType:
			if specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek == nil {
				specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek = []*SpecificationSeek{newSpecificationSeek()}
			} else {
				specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek = append(specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek, newSpecificationSeek())
			}
			return &specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek[len(specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek)-1].SpecificationElement, true
		case SeekIDIdType:
			if specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek[len(specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek)-1].SeekID == nil {
				specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek[len(specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek)-1].SeekID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek[len(specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek)-1].SeekID, false
		case SeekPositionIdType:
			if specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek[len(specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek)-1].SeekPosition == nil {
				specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek[len(specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek)-1].SeekPosition = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek[len(specification.Segment[len(specification.Segment)-1].SeekHead[len(specification.Segment[len(specification.Segment)-1].SeekHead)-1].Seek)-1].SeekPosition, false
		case InfoIdType:
			if specification.Segment[len(specification.Segment)-1].Info == nil {
				specification.Segment[len(specification.Segment)-1].Info = []*SpecificationInfo{newSpecificationInfo()}
			} else {
				specification.Segment[len(specification.Segment)-1].Info = append(specification.Segment[len(specification.Segment)-1].Info, newSpecificationInfo())
			}
			return &specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SpecificationElement, true
		case SegmentUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentUID == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentUID, false
		case SegmentFilenameIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentFilename == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentFilename = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentFilename, false
		case PrevUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].PrevUID == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].PrevUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].PrevUID, false
		case PrevFilenameIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].PrevFilename == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].PrevFilename = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].PrevFilename, false
		case NextUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].NextUID == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].NextUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].NextUID, false
		case NextFilenameIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].NextFilename == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].NextFilename = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].NextFilename, false
		case SegmentFamilyIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentFamily == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentFamily = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentFamily = append(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentFamily, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentFamily[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].SegmentFamily)-1], false
		case ChapterTranslateIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate = []*SpecificationChapterTranslate{newSpecificationChapterTranslate()}
			} else {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate = append(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate, newSpecificationChapterTranslate())
			}
			return &specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].SpecificationElement, true
		case ChapterTranslateEditionUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateEditionUID == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateEditionUID = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateEditionUID = append(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateEditionUID, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateEditionUID[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateEditionUID)-1], false
		case ChapterTranslateCodecIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateCodec == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateCodec = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateCodec, false
		case ChapterTranslateIDIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateID == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate[len(specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].ChapterTranslate)-1].ChapterTranslateID, false
		case TimecodeScaleIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].TimecodeScale == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].TimecodeScale = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].TimecodeScale, false
		case DurationIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].Duration == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].Duration = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].Duration, false
		case DateUTCIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].DateUTC == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].DateUTC = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].DateUTC, false
		case TitleIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].Title == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].Title = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].Title, false
		case MuxingAppIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].MuxingApp == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].MuxingApp = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].MuxingApp, false
		case WritingAppIdType:
			if specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].WritingApp == nil {
				specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].WritingApp = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Info[len(specification.Segment[len(specification.Segment)-1].Info)-1].WritingApp, false
		case ClusterIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster == nil {
				specification.Segment[len(specification.Segment)-1].Cluster = []*SpecificationCluster{newSpecificationCluster()}
			} else {
				specification.Segment[len(specification.Segment)-1].Cluster = append(specification.Segment[len(specification.Segment)-1].Cluster, newSpecificationCluster())
			}
			return &specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SpecificationElement, true
		case TimecodeIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].Timecode == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].Timecode = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].Timecode, false
		case SilentTracksIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SilentTracks == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SilentTracks = newSpecificationSilentTracks()
			}
			return &specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SilentTracks.SpecificationElement, true
		case SilentTrackNumberIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SilentTracks.SilentTrackNumber == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SilentTracks.SilentTrackNumber = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SilentTracks.SilentTrackNumber = append(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SilentTracks.SilentTrackNumber, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SilentTracks.SilentTrackNumber[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SilentTracks.SilentTrackNumber)-1], false
		case PositionIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].Position == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].Position = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].Position, false
		case PrevSizeIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].PrevSize == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].PrevSize = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].PrevSize, false
		case SimpleBlockIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SimpleBlock == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SimpleBlock = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SimpleBlock = append(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SimpleBlock, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SimpleBlock[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].SimpleBlock)-1], false
		case BlockGroupIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup = []*SpecificationBlockGroup{newSpecificationBlockGroup()}
			} else {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup = append(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup, newSpecificationBlockGroup())
			}
			return &specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].SpecificationElement, true
		case BlockIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Block == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Block = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Block, false
		case BlockVirtualIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockVirtual == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockVirtual = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockVirtual, false
		case BlockAdditionsIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions = newSpecificationBlockAdditions()
			}
			return &specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.SpecificationElement, true
		case BlockMoreIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore = []*SpecificationBlockMore{newSpecificationBlockMore()}
			} else {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore = append(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore, newSpecificationBlockMore())
			}
			return &specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore)-1].SpecificationElement, true
		case BlockAddIDIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore)-1].BlockAddID == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore)-1].BlockAddID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore)-1].BlockAddID, false
		case BlockAdditionalIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore)-1].BlockAdditional == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore)-1].BlockAdditional = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockAdditions.BlockMore)-1].BlockAdditional, false
		case BlockDurationIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockDuration == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockDuration = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].BlockDuration, false
		case ReferencePriorityIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferencePriority == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferencePriority = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferencePriority, false
		case ReferenceBlockIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceBlock == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceBlock = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceBlock = append(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceBlock, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceBlock[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceBlock)-1], false
		case ReferenceVirtualIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceVirtual == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceVirtual = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceVirtual, false
		case CodecStateIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].CodecState == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].CodecState = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].CodecState, false
		case DiscardPaddingIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].DiscardPadding == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].DiscardPadding = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].DiscardPadding, false
		case SlicesIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices = newSpecificationSlices()
			}
			return &specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.SpecificationElement, true
		case TimeSliceIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice = []*SpecificationTimeSlice{newSpecificationTimeSlice()}
			} else {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice = append(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice, newSpecificationTimeSlice())
			}
			return &specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].SpecificationElement, true
		case LaceNumberIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].LaceNumber == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].LaceNumber = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].LaceNumber, false
		case FrameNumberIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].FrameNumber == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].FrameNumber = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].FrameNumber, false
		case BlockAdditionIDIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].BlockAdditionID == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].BlockAdditionID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].BlockAdditionID, false
		case DelayIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].Delay == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].Delay = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].Delay, false
		case SliceDurationIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].SliceDuration == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].SliceDuration = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].Slices.TimeSlice)-1].SliceDuration, false
		case ReferenceFrameIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceFrame == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceFrame = newSpecificationReferenceFrame()
			}
			return &specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceFrame.SpecificationElement, true
		case ReferenceOffsetIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceFrame.ReferenceOffset == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceFrame.ReferenceOffset = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceFrame.ReferenceOffset, false
		case ReferenceTimeCodeIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceFrame.ReferenceTimeCode == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceFrame.ReferenceTimeCode = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].BlockGroup)-1].ReferenceFrame.ReferenceTimeCode, false
		case EncryptedBlockIdType:
			if specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].EncryptedBlock == nil {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].EncryptedBlock = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].EncryptedBlock = append(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].EncryptedBlock, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].EncryptedBlock[len(specification.Segment[len(specification.Segment)-1].Cluster[len(specification.Segment[len(specification.Segment)-1].Cluster)-1].EncryptedBlock)-1], false
		case TracksIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks == nil {
				specification.Segment[len(specification.Segment)-1].Tracks = []*SpecificationTracks{newSpecificationTracks()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tracks = append(specification.Segment[len(specification.Segment)-1].Tracks, newSpecificationTracks())
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].SpecificationElement, true
		case TrackEntryIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry = []*SpecificationTrackEntry{newSpecificationTrackEntry()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry = append(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry, newSpecificationTrackEntry())
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].SpecificationElement, true
		case TrackNumberIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackNumber == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackNumber = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackNumber, false
		case TrackUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackUID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackUID, false
		case TrackTypeIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackType == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackType = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackType, false
		case FlagEnabledIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagEnabled == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagEnabled = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagEnabled, false
		case FlagDefaultIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagDefault == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagDefault = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagDefault, false
		case FlagForcedIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagForced == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagForced = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagForced, false
		case FlagLacingIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagLacing == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagLacing = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].FlagLacing, false
		case MinCacheIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].MinCache == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].MinCache = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].MinCache, false
		case MaxCacheIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].MaxCache == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].MaxCache = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].MaxCache, false
		case DefaultDurationIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].DefaultDuration == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].DefaultDuration = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].DefaultDuration, false
		case DefaultDecodedFieldDurationIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].DefaultDecodedFieldDuration == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].DefaultDecodedFieldDuration = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].DefaultDecodedFieldDuration, false
		case TrackTimecodeScaleIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTimecodeScale == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTimecodeScale = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTimecodeScale, false
		case TrackOffsetIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOffset == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOffset = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOffset, false
		case MaxBlockAdditionIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].MaxBlockAdditionID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].MaxBlockAdditionID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].MaxBlockAdditionID, false
		case NameIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Name == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Name = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Name, false
		case LanguageIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Language == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Language = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Language, false
		case CodecIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecID, false
		case CodecPrivateIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecPrivate == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecPrivate = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecPrivate, false
		case CodecNameIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecName == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecName = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecName, false
		case AttachmentLinkIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].AttachmentLink == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].AttachmentLink = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].AttachmentLink, false
		case CodecSettingsIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecSettings == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecSettings = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecSettings, false
		case CodecInfoURLIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecInfoURL == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecInfoURL = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecInfoURL = append(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecInfoURL, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecInfoURL[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecInfoURL)-1], false
		case CodecDownloadURLIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDownloadURL == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDownloadURL = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDownloadURL = append(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDownloadURL, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDownloadURL[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDownloadURL)-1], false
		case CodecDecodeAllIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDecodeAll == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDecodeAll = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDecodeAll, false
		case TrackOverlayIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOverlay == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOverlay = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOverlay = append(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOverlay, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOverlay[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOverlay)-1], false
		case CodecDelayIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDelay == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDelay = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].CodecDelay, false
		case SeekPreRollIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].SeekPreRoll == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].SeekPreRoll = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].SeekPreRoll, false
		case TrackTranslateIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate = []*SpecificationTrackTranslate{newSpecificationTrackTranslate()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate = append(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate, newSpecificationTrackTranslate())
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].SpecificationElement, true
		case TrackTranslateEditionUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateEditionUID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateEditionUID = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateEditionUID = append(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateEditionUID, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateEditionUID[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateEditionUID)-1], false
		case TrackTranslateCodecIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateCodec == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateCodec = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateCodec, false
		case TrackTranslateTrackIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateTrackID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateTrackID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackTranslate)-1].TrackTranslateTrackID, false
		case VideoIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video = newSpecificationVideo()
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.SpecificationElement, true
		case FlagInterlacedIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.FlagInterlaced == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.FlagInterlaced = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.FlagInterlaced, false
		case StereoModeIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.StereoMode == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.StereoMode = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.StereoMode, false
		case AlphaModeIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.AlphaMode == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.AlphaMode = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.AlphaMode, false
		case OldStereoModeIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.OldStereoMode == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.OldStereoMode = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.OldStereoMode, false
		case PixelWidthIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelWidth == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelWidth = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelWidth, false
		case PixelHeightIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelHeight == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelHeight = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelHeight, false
		case PixelCropBottomIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropBottom == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropBottom = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropBottom, false
		case PixelCropTopIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropTop == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropTop = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropTop, false
		case PixelCropLeftIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropLeft == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropLeft = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropLeft, false
		case PixelCropRightIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropRight == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropRight = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.PixelCropRight, false
		case DisplayWidthIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.DisplayWidth == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.DisplayWidth = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.DisplayWidth, false
		case DisplayHeightIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.DisplayHeight == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.DisplayHeight = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.DisplayHeight, false
		case DisplayUnitIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.DisplayUnit == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.DisplayUnit = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.DisplayUnit, false
		case AspectRatioTypeIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.AspectRatioType == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.AspectRatioType = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.AspectRatioType, false
		case ColourSpaceIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.ColourSpace == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.ColourSpace = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.ColourSpace, false
		case GammaValueIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.GammaValue == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.GammaValue = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.GammaValue, false
		case FrameRateIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.FrameRate == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.FrameRate = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Video.FrameRate, false
		case AudioIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio = newSpecificationAudio()
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.SpecificationElement, true
		case SamplingFrequencyIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.SamplingFrequency == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.SamplingFrequency = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.SamplingFrequency, false
		case OutputSamplingFrequencyIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.OutputSamplingFrequency == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.OutputSamplingFrequency = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.OutputSamplingFrequency, false
		case ChannelsIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.Channels == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.Channels = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.Channels, false
		case ChannelPositionsIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.ChannelPositions == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.ChannelPositions = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.ChannelPositions, false
		case BitDepthIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.BitDepth == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.BitDepth = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].Audio.BitDepth, false
		case TrackOperationIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation = newSpecificationTrackOperation()
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.SpecificationElement, true
		case TrackCombinePlanesIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes = newSpecificationTrackCombinePlanes()
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.SpecificationElement, true
		case TrackPlaneIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane = []*SpecificationTrackPlane{newSpecificationTrackPlane()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane = append(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane, newSpecificationTrackPlane())
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane)-1].SpecificationElement, true
		case TrackPlaneUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane)-1].TrackPlaneUID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane)-1].TrackPlaneUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane)-1].TrackPlaneUID, false
		case TrackPlaneTypeIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane)-1].TrackPlaneType == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane)-1].TrackPlaneType = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackCombinePlanes.TrackPlane)-1].TrackPlaneType, false
		case TrackJoinBlocksIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackJoinBlocks == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackJoinBlocks = newSpecificationTrackJoinBlocks()
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackJoinBlocks.SpecificationElement, true
		case TrackJoinUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackJoinBlocks.TrackJoinUID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackJoinBlocks.TrackJoinUID = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackJoinBlocks.TrackJoinUID = append(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackJoinBlocks.TrackJoinUID, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackJoinBlocks.TrackJoinUID[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrackOperation.TrackJoinBlocks.TrackJoinUID)-1], false
		case TrickTrackUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickTrackUID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickTrackUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickTrackUID, false
		case TrickTrackSegmentUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickTrackSegmentUID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickTrackSegmentUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickTrackSegmentUID, false
		case TrickTrackFlagIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickTrackFlag == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickTrackFlag = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickTrackFlag, false
		case TrickMasterTrackUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickMasterTrackUID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickMasterTrackUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickMasterTrackUID, false
		case TrickMasterTrackSegmentUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickMasterTrackSegmentUID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickMasterTrackSegmentUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].TrickMasterTrackSegmentUID, false
		case ContentEncodingsIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings = newSpecificationContentEncodings()
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.SpecificationElement, true
		case ContentEncodingIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding = []*SpecificationContentEncoding{newSpecificationContentEncoding()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding = append(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding, newSpecificationContentEncoding())
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].SpecificationElement, true
		case ContentEncodingOrderIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncodingOrder == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncodingOrder = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncodingOrder, false
		case ContentEncodingScopeIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncodingScope == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncodingScope = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncodingScope, false
		case ContentEncodingTypeIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncodingType == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncodingType = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncodingType, false
		case ContentCompressionIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentCompression == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentCompression = newSpecificationContentCompression()
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentCompression.SpecificationElement, true
		case ContentCompAlgoIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentCompression.ContentCompAlgo == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentCompression.ContentCompAlgo = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentCompression.ContentCompAlgo, false
		case ContentCompSettingsIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentCompression.ContentCompSettings == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentCompression.ContentCompSettings = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentCompression.ContentCompSettings, false
		case ContentEncryptionIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption = newSpecificationContentEncryption()
			}
			return &specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.SpecificationElement, true
		case ContentEncAlgoIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentEncAlgo == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentEncAlgo = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentEncAlgo, false
		case ContentEncKeyIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentEncKeyID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentEncKeyID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentEncKeyID, false
		case ContentSignatureIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSignature == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSignature = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSignature, false
		case ContentSigKeyIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSigKeyID == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSigKeyID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSigKeyID, false
		case ContentSigAlgoIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSigAlgo == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSigAlgo = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSigAlgo, false
		case ContentSigHashAlgoIdType:
			if specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSigHashAlgo == nil {
				specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSigHashAlgo = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry[len(specification.Segment[len(specification.Segment)-1].Tracks[len(specification.Segment[len(specification.Segment)-1].Tracks)-1].TrackEntry)-1].ContentEncodings.ContentEncoding)-1].ContentEncryption.ContentSigHashAlgo, false
		case CuesIdType:
			if specification.Segment[len(specification.Segment)-1].Cues == nil {
				specification.Segment[len(specification.Segment)-1].Cues = newSpecificationCues()
			}
			return &specification.Segment[len(specification.Segment)-1].Cues.SpecificationElement, true
		case CuePointIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint = []*SpecificationCuePoint{newSpecificationCuePoint()}
			} else {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint = append(specification.Segment[len(specification.Segment)-1].Cues.CuePoint, newSpecificationCuePoint())
			}
			return &specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].SpecificationElement, true
		case CueTimeIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTime == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTime = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTime, false
		case CueTrackPositionsIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions = []*SpecificationCueTrackPositions{newSpecificationCueTrackPositions()}
			} else {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions = append(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions, newSpecificationCueTrackPositions())
			}
			return &specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].SpecificationElement, true
		case CueTrackIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueTrack == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueTrack = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueTrack, false
		case CueClusterPositionIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueClusterPosition == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueClusterPosition = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueClusterPosition, false
		case CueRelativePositionIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueRelativePosition == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueRelativePosition = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueRelativePosition, false
		case CueDurationIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueDuration == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueDuration = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueDuration, false
		case CueBlockNumberIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueBlockNumber == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueBlockNumber = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueBlockNumber, false
		case CueCodecStateIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueCodecState == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueCodecState = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueCodecState, false
		case CueReferenceIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference = []*SpecificationCueReference{newSpecificationCueReference()}
			} else {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference = append(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference, newSpecificationCueReference())
			}
			return &specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].SpecificationElement, true
		case CueRefTimeIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefTime == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefTime = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefTime, false
		case CueRefClusterIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefCluster == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefCluster = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefCluster, false
		case CueRefNumberIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefNumber == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefNumber = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefNumber, false
		case CueRefCodecStateIdType:
			if specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefCodecState == nil {
				specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefCodecState = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint[len(specification.Segment[len(specification.Segment)-1].Cues.CuePoint)-1].CueTrackPositions)-1].CueReference)-1].CueRefCodecState, false
		case AttachmentsIdType:
			if specification.Segment[len(specification.Segment)-1].Attachments == nil {
				specification.Segment[len(specification.Segment)-1].Attachments = newSpecificationAttachments()
			}
			return &specification.Segment[len(specification.Segment)-1].Attachments.SpecificationElement, true
		case AttachedFileIdType:
			if specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile == nil {
				specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile = []*SpecificationAttachedFile{newSpecificationAttachedFile()}
			} else {
				specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile = append(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile, newSpecificationAttachedFile())
			}
			return &specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].SpecificationElement, true
		case FileDescriptionIdType:
			if specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileDescription == nil {
				specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileDescription = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileDescription, false
		case FileNameIdType:
			if specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileName == nil {
				specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileName = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileName, false
		case FileMimeTypeIdType:
			if specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileMimeType == nil {
				specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileMimeType = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileMimeType, false
		case FileDataIdType:
			if specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileData == nil {
				specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileData = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileData, false
		case FileUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileUID == nil {
				specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileUID, false
		case FileReferralIdType:
			if specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileReferral == nil {
				specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileReferral = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileReferral, false
		case FileUsedStartTimeIdType:
			if specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileUsedStartTime == nil {
				specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileUsedStartTime = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileUsedStartTime, false
		case FileUsedEndTimeIdType:
			if specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileUsedEndTime == nil {
				specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileUsedEndTime = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile[len(specification.Segment[len(specification.Segment)-1].Attachments.AttachedFile)-1].FileUsedEndTime, false
		case ChaptersIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters == nil {
				specification.Segment[len(specification.Segment)-1].Chapters = newSpecificationChapters()
			}
			return &specification.Segment[len(specification.Segment)-1].Chapters.SpecificationElement, true
		case EditionEntryIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry = []*SpecificationEditionEntry{newSpecificationEditionEntry()}
			} else {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry = append(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry, newSpecificationEditionEntry())
			}
			return &specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].SpecificationElement, true
		case EditionUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionUID == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionUID, false
		case EditionFlagHiddenIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionFlagHidden == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionFlagHidden = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionFlagHidden, false
		case EditionFlagDefaultIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionFlagDefault == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionFlagDefault = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionFlagDefault, false
		case EditionFlagOrderedIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionFlagOrdered == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionFlagOrdered = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].EditionFlagOrdered, false
		case ChapterAtomIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom = []*SpecificationChapterAtom{newSpecificationChapterAtom()}
			} else {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom = append(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom, newSpecificationChapterAtom())
			}
			return &specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].SpecificationElement, true
		case ChapterUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterUID == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterUID, false
		case ChapterStringUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterStringUID == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterStringUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterStringUID, false
		case ChapterTimeStartIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTimeStart == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTimeStart = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTimeStart, false
		case ChapterTimeEndIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTimeEnd == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTimeEnd = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTimeEnd, false
		case ChapterFlagHiddenIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterFlagHidden == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterFlagHidden = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterFlagHidden, false
		case ChapterFlagEnabledIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterFlagEnabled == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterFlagEnabled = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterFlagEnabled, false
		case ChapterSegmentUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterSegmentUID == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterSegmentUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterSegmentUID, false
		case ChapterSegmentEditionUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterSegmentEditionUID == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterSegmentEditionUID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterSegmentEditionUID, false
		case ChapterPhysicalEquivIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterPhysicalEquiv == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterPhysicalEquiv = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterPhysicalEquiv, false
		case ChapterTrackIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTrack == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTrack = newSpecificationChapterTrack()
			}
			return &specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTrack.SpecificationElement, true
		case ChapterTrackNumberIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTrack.ChapterTrackNumber == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTrack.ChapterTrackNumber = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTrack.ChapterTrackNumber = append(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTrack.ChapterTrackNumber, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTrack.ChapterTrackNumber[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterTrack.ChapterTrackNumber)-1], false
		case ChapterDisplayIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay = []*SpecificationChapterDisplay{newSpecificationChapterDisplay()}
			} else {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay = append(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay, newSpecificationChapterDisplay())
			}
			return &specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].SpecificationElement, true
		case ChapStringIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapString == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapString = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapString, false
		case ChapLanguageIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapLanguage == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapLanguage = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapLanguage = append(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapLanguage, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapLanguage[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapLanguage)-1], false
		case ChapCountryIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapCountry == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapCountry = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapCountry = append(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapCountry, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapCountry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapterDisplay)-1].ChapCountry)-1], false
		case ChapProcessIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess = []*SpecificationChapProcess{newSpecificationChapProcess()}
			} else {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess = append(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess, newSpecificationChapProcess())
			}
			return &specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].SpecificationElement, true
		case ChapProcessCodecIDIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCodecID == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCodecID = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCodecID, false
		case ChapProcessPrivateIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessPrivate == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessPrivate = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessPrivate, false
		case ChapProcessCommandIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand = []*SpecificationChapProcessCommand{newSpecificationChapProcessCommand()}
			} else {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand = append(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand, newSpecificationChapProcessCommand())
			}
			return &specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand)-1].SpecificationElement, true
		case ChapProcessTimeIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand)-1].ChapProcessTime == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand)-1].ChapProcessTime = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand)-1].ChapProcessTime, false
		case ChapProcessDataIdType:
			if specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand)-1].ChapProcessData == nil {
				specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand)-1].ChapProcessData = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry[len(specification.Segment[len(specification.Segment)-1].Chapters.EditionEntry)-1].ChapterAtom)-1].ChapProcess)-1].ChapProcessCommand)-1].ChapProcessData, false
		case TagsIdType:
			if specification.Segment[len(specification.Segment)-1].Tags == nil {
				specification.Segment[len(specification.Segment)-1].Tags = []*SpecificationTags{newSpecificationTags()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tags = append(specification.Segment[len(specification.Segment)-1].Tags, newSpecificationTags())
			}
			return &specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].SpecificationElement, true
		case TagIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag = []*SpecificationTag{newSpecificationTag()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag = append(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag, newSpecificationTag())
			}
			return &specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SpecificationElement, true
		case TargetsIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets = newSpecificationTargets()
			}
			return &specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.SpecificationElement, true
		case TargetTypeValueIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TargetTypeValue == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TargetTypeValue = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TargetTypeValue, false
		case TargetTypeIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TargetType == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TargetType = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TargetType, false
		case TagTrackUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagTrackUID == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagTrackUID = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagTrackUID = append(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagTrackUID, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagTrackUID[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagTrackUID)-1], false
		case TagEditionUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagEditionUID == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagEditionUID = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagEditionUID = append(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagEditionUID, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagEditionUID[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagEditionUID)-1], false
		case TagChapterUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagChapterUID == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagChapterUID = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagChapterUID = append(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagChapterUID, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagChapterUID[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagChapterUID)-1], false
		case TagAttachmentUIDIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagAttachmentUID == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagAttachmentUID = []*SpecificationElement{newSpecificationElement()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagAttachmentUID = append(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagAttachmentUID, newSpecificationElement())
			}
			return specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagAttachmentUID[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].Targets.TagAttachmentUID)-1], false
		case SimpleTagIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag = []*SpecificationSimpleTag{newSpecificationSimpleTag()}
			} else {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag = append(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag, newSpecificationSimpleTag())
			}
			return &specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].SpecificationElement, true
		case TagNameIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagName == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagName = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagName, false
		case TagLanguageIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagLanguage == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagLanguage = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagLanguage, false
		case TagDefaultIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagDefault == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagDefault = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagDefault, false
		case TagStringIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagString == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagString = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagString, false
		case TagBinaryIdType:
			if specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagBinary == nil {
				specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagBinary = newSpecificationElement()
			}
			return specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag[len(specification.Segment[len(specification.Segment)-1].Tags[len(specification.Segment[len(specification.Segment)-1].Tags)-1].Tag)-1].SimpleTag)-1].TagBinary, false
		default:
			return 
	}
}

