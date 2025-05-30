// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/scorer.proto

package scorer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScoreRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StartDate     string                 `protobuf:"bytes,1,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate       string                 `protobuf:"bytes,2,opt,name=endDate,proto3" json:"endDate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScoreRequest) Reset() {
	*x = ScoreRequest{}
	mi := &file_proto_scorer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreRequest) ProtoMessage() {}

func (x *ScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scorer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreRequest.ProtoReflect.Descriptor instead.
func (*ScoreRequest) Descriptor() ([]byte, []int) {
	return file_proto_scorer_proto_rawDescGZIP(), []int{0}
}

func (x *ScoreRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *ScoreRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type OverallScoreResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ScorePercentage float64                `protobuf:"fixed64,1,opt,name=scorePercentage,proto3" json:"scorePercentage,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *OverallScoreResponse) Reset() {
	*x = OverallScoreResponse{}
	mi := &file_proto_scorer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverallScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverallScoreResponse) ProtoMessage() {}

func (x *OverallScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scorer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverallScoreResponse.ProtoReflect.Descriptor instead.
func (*OverallScoreResponse) Descriptor() ([]byte, []int) {
	return file_proto_scorer_proto_rawDescGZIP(), []int{1}
}

func (x *OverallScoreResponse) GetScorePercentage() float64 {
	if x != nil {
		return x.ScorePercentage
	}
	return 0
}

type TicketScore struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	TicketId       int64                  `protobuf:"varint,1,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
	CategoryScores map[string]float64     `protobuf:"bytes,2,rep,name=categoryScores,proto3" json:"categoryScores,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TicketScore) Reset() {
	*x = TicketScore{}
	mi := &file_proto_scorer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketScore) ProtoMessage() {}

func (x *TicketScore) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scorer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketScore.ProtoReflect.Descriptor instead.
func (*TicketScore) Descriptor() ([]byte, []int) {
	return file_proto_scorer_proto_rawDescGZIP(), []int{2}
}

func (x *TicketScore) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

func (x *TicketScore) GetCategoryScores() map[string]float64 {
	if x != nil {
		return x.CategoryScores
	}
	return nil
}

type CategoryScoresByTicketResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TicketScores  []*TicketScore         `protobuf:"bytes,1,rep,name=ticketScores,proto3" json:"ticketScores,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryScoresByTicketResponse) Reset() {
	*x = CategoryScoresByTicketResponse{}
	mi := &file_proto_scorer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryScoresByTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryScoresByTicketResponse) ProtoMessage() {}

func (x *CategoryScoresByTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scorer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryScoresByTicketResponse.ProtoReflect.Descriptor instead.
func (*CategoryScoresByTicketResponse) Descriptor() ([]byte, []int) {
	return file_proto_scorer_proto_rawDescGZIP(), []int{3}
}

func (x *CategoryScoresByTicketResponse) GetTicketScores() []*TicketScore {
	if x != nil {
		return x.TicketScores
	}
	return nil
}

type CategoryScore struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Category      string                 `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	RatingCount   int32                  `protobuf:"varint,2,opt,name=ratingCount,proto3" json:"ratingCount,omitempty"`
	DateToScore   []*DateScore           `protobuf:"bytes,3,rep,name=dateToScore,proto3" json:"dateToScore,omitempty"`
	OverallScore  float64                `protobuf:"fixed64,4,opt,name=overallScore,proto3" json:"overallScore,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryScore) Reset() {
	*x = CategoryScore{}
	mi := &file_proto_scorer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryScore) ProtoMessage() {}

func (x *CategoryScore) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scorer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryScore.ProtoReflect.Descriptor instead.
func (*CategoryScore) Descriptor() ([]byte, []int) {
	return file_proto_scorer_proto_rawDescGZIP(), []int{4}
}

func (x *CategoryScore) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CategoryScore) GetRatingCount() int32 {
	if x != nil {
		return x.RatingCount
	}
	return 0
}

func (x *CategoryScore) GetDateToScore() []*DateScore {
	if x != nil {
		return x.DateToScore
	}
	return nil
}

func (x *CategoryScore) GetOverallScore() float64 {
	if x != nil {
		return x.OverallScore
	}
	return 0
}

type DateScore struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Date          string                 `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Score         float64                `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DateScore) Reset() {
	*x = DateScore{}
	mi := &file_proto_scorer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DateScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateScore) ProtoMessage() {}

func (x *DateScore) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scorer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateScore.ProtoReflect.Descriptor instead.
func (*DateScore) Descriptor() ([]byte, []int) {
	return file_proto_scorer_proto_rawDescGZIP(), []int{5}
}

func (x *DateScore) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *DateScore) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type CategoryScoresResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Categories    []*CategoryScore       `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryScoresResponse) Reset() {
	*x = CategoryScoresResponse{}
	mi := &file_proto_scorer_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryScoresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryScoresResponse) ProtoMessage() {}

func (x *CategoryScoresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scorer_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryScoresResponse.ProtoReflect.Descriptor instead.
func (*CategoryScoresResponse) Descriptor() ([]byte, []int) {
	return file_proto_scorer_proto_rawDescGZIP(), []int{6}
}

func (x *CategoryScoresResponse) GetCategories() []*CategoryScore {
	if x != nil {
		return x.Categories
	}
	return nil
}

type PeriodOverPeriodChangeResponse struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	PreviousPeriodStart string                 `protobuf:"bytes,1,opt,name=previousPeriodStart,proto3" json:"previousPeriodStart,omitempty"`
	PreviousPeriodEnd   string                 `protobuf:"bytes,2,opt,name=previousPeriodEnd,proto3" json:"previousPeriodEnd,omitempty"`
	PreviousPeriodScore float64                `protobuf:"fixed64,3,opt,name=previousPeriodScore,proto3" json:"previousPeriodScore,omitempty"`
	CurrentPeriodStart  string                 `protobuf:"bytes,4,opt,name=currentPeriodStart,proto3" json:"currentPeriodStart,omitempty"`
	CurrentPeriodEnd    string                 `protobuf:"bytes,5,opt,name=currentPeriodEnd,proto3" json:"currentPeriodEnd,omitempty"`
	CurrentPeriodScore  float64                `protobuf:"fixed64,6,opt,name=currentPeriodScore,proto3" json:"currentPeriodScore,omitempty"`
	ChangePercentage    float64                `protobuf:"fixed64,7,opt,name=changePercentage,proto3" json:"changePercentage,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *PeriodOverPeriodChangeResponse) Reset() {
	*x = PeriodOverPeriodChangeResponse{}
	mi := &file_proto_scorer_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PeriodOverPeriodChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeriodOverPeriodChangeResponse) ProtoMessage() {}

func (x *PeriodOverPeriodChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scorer_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeriodOverPeriodChangeResponse.ProtoReflect.Descriptor instead.
func (*PeriodOverPeriodChangeResponse) Descriptor() ([]byte, []int) {
	return file_proto_scorer_proto_rawDescGZIP(), []int{7}
}

func (x *PeriodOverPeriodChangeResponse) GetPreviousPeriodStart() string {
	if x != nil {
		return x.PreviousPeriodStart
	}
	return ""
}

func (x *PeriodOverPeriodChangeResponse) GetPreviousPeriodEnd() string {
	if x != nil {
		return x.PreviousPeriodEnd
	}
	return ""
}

func (x *PeriodOverPeriodChangeResponse) GetPreviousPeriodScore() float64 {
	if x != nil {
		return x.PreviousPeriodScore
	}
	return 0
}

func (x *PeriodOverPeriodChangeResponse) GetCurrentPeriodStart() string {
	if x != nil {
		return x.CurrentPeriodStart
	}
	return ""
}

func (x *PeriodOverPeriodChangeResponse) GetCurrentPeriodEnd() string {
	if x != nil {
		return x.CurrentPeriodEnd
	}
	return ""
}

func (x *PeriodOverPeriodChangeResponse) GetCurrentPeriodScore() float64 {
	if x != nil {
		return x.CurrentPeriodScore
	}
	return 0
}

func (x *PeriodOverPeriodChangeResponse) GetChangePercentage() float64 {
	if x != nil {
		return x.ChangePercentage
	}
	return 0
}

var File_proto_scorer_proto protoreflect.FileDescriptor

const file_proto_scorer_proto_rawDesc = "" +
	"\n" +
	"\x12proto/scorer.proto\x12\x06scorer\"F\n" +
	"\fScoreRequest\x12\x1c\n" +
	"\tstartDate\x18\x01 \x01(\tR\tstartDate\x12\x18\n" +
	"\aendDate\x18\x02 \x01(\tR\aendDate\"@\n" +
	"\x14OverallScoreResponse\x12(\n" +
	"\x0fscorePercentage\x18\x01 \x01(\x01R\x0fscorePercentage\"\xbd\x01\n" +
	"\vTicketScore\x12\x1a\n" +
	"\bticketId\x18\x01 \x01(\x03R\bticketId\x12O\n" +
	"\x0ecategoryScores\x18\x02 \x03(\v2'.scorer.TicketScore.CategoryScoresEntryR\x0ecategoryScores\x1aA\n" +
	"\x13CategoryScoresEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x01R\x05value:\x028\x01\"Y\n" +
	"\x1eCategoryScoresByTicketResponse\x127\n" +
	"\fticketScores\x18\x01 \x03(\v2\x13.scorer.TicketScoreR\fticketScores\"\xa6\x01\n" +
	"\rCategoryScore\x12\x1a\n" +
	"\bcategory\x18\x01 \x01(\tR\bcategory\x12 \n" +
	"\vratingCount\x18\x02 \x01(\x05R\vratingCount\x123\n" +
	"\vdateToScore\x18\x03 \x03(\v2\x11.scorer.DateScoreR\vdateToScore\x12\"\n" +
	"\foverallScore\x18\x04 \x01(\x01R\foverallScore\"5\n" +
	"\tDateScore\x12\x12\n" +
	"\x04date\x18\x01 \x01(\tR\x04date\x12\x14\n" +
	"\x05score\x18\x02 \x01(\x01R\x05score\"O\n" +
	"\x16CategoryScoresResponse\x125\n" +
	"\n" +
	"categories\x18\x01 \x03(\v2\x15.scorer.CategoryScoreR\n" +
	"categories\"\xea\x02\n" +
	"\x1ePeriodOverPeriodChangeResponse\x120\n" +
	"\x13previousPeriodStart\x18\x01 \x01(\tR\x13previousPeriodStart\x12,\n" +
	"\x11previousPeriodEnd\x18\x02 \x01(\tR\x11previousPeriodEnd\x120\n" +
	"\x13previousPeriodScore\x18\x03 \x01(\x01R\x13previousPeriodScore\x12.\n" +
	"\x12currentPeriodStart\x18\x04 \x01(\tR\x12currentPeriodStart\x12*\n" +
	"\x10currentPeriodEnd\x18\x05 \x01(\tR\x10currentPeriodEnd\x12.\n" +
	"\x12currentPeriodScore\x18\x06 \x01(\x01R\x12currentPeriodScore\x12*\n" +
	"\x10changePercentage\x18\a \x01(\x01R\x10changePercentage2\xd7\x02\n" +
	"\rScorerService\x12I\n" +
	"\x11GetCategoryScores\x12\x14.scorer.ScoreRequest\x1a\x1e.scorer.CategoryScoresResponse\x12Y\n" +
	"\x19GetCategoryScoresByTicket\x12\x14.scorer.ScoreRequest\x1a&.scorer.CategoryScoresByTicketResponse\x12E\n" +
	"\x0fGetOverallScore\x12\x14.scorer.ScoreRequest\x1a\x1c.scorer.OverallScoreResponse\x12Y\n" +
	"\x19GetPeriodOverPeriodChange\x12\x14.scorer.ScoreRequest\x1a&.scorer.PeriodOverPeriodChangeResponseB3Z1github.com/siimsams/zendesk-homework/proto;scorerb\x06proto3"

var (
	file_proto_scorer_proto_rawDescOnce sync.Once
	file_proto_scorer_proto_rawDescData []byte
)

func file_proto_scorer_proto_rawDescGZIP() []byte {
	file_proto_scorer_proto_rawDescOnce.Do(func() {
		file_proto_scorer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scorer_proto_rawDesc), len(file_proto_scorer_proto_rawDesc)))
	})
	return file_proto_scorer_proto_rawDescData
}

var file_proto_scorer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_scorer_proto_goTypes = []any{
	(*ScoreRequest)(nil),                   // 0: scorer.ScoreRequest
	(*OverallScoreResponse)(nil),           // 1: scorer.OverallScoreResponse
	(*TicketScore)(nil),                    // 2: scorer.TicketScore
	(*CategoryScoresByTicketResponse)(nil), // 3: scorer.CategoryScoresByTicketResponse
	(*CategoryScore)(nil),                  // 4: scorer.CategoryScore
	(*DateScore)(nil),                      // 5: scorer.DateScore
	(*CategoryScoresResponse)(nil),         // 6: scorer.CategoryScoresResponse
	(*PeriodOverPeriodChangeResponse)(nil), // 7: scorer.PeriodOverPeriodChangeResponse
	nil,                                    // 8: scorer.TicketScore.CategoryScoresEntry
}
var file_proto_scorer_proto_depIdxs = []int32{
	8, // 0: scorer.TicketScore.categoryScores:type_name -> scorer.TicketScore.CategoryScoresEntry
	2, // 1: scorer.CategoryScoresByTicketResponse.ticketScores:type_name -> scorer.TicketScore
	5, // 2: scorer.CategoryScore.dateToScore:type_name -> scorer.DateScore
	4, // 3: scorer.CategoryScoresResponse.categories:type_name -> scorer.CategoryScore
	0, // 4: scorer.ScorerService.GetCategoryScores:input_type -> scorer.ScoreRequest
	0, // 5: scorer.ScorerService.GetCategoryScoresByTicket:input_type -> scorer.ScoreRequest
	0, // 6: scorer.ScorerService.GetOverallScore:input_type -> scorer.ScoreRequest
	0, // 7: scorer.ScorerService.GetPeriodOverPeriodChange:input_type -> scorer.ScoreRequest
	6, // 8: scorer.ScorerService.GetCategoryScores:output_type -> scorer.CategoryScoresResponse
	3, // 9: scorer.ScorerService.GetCategoryScoresByTicket:output_type -> scorer.CategoryScoresByTicketResponse
	1, // 10: scorer.ScorerService.GetOverallScore:output_type -> scorer.OverallScoreResponse
	7, // 11: scorer.ScorerService.GetPeriodOverPeriodChange:output_type -> scorer.PeriodOverPeriodChangeResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_scorer_proto_init() }
func file_proto_scorer_proto_init() {
	if File_proto_scorer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scorer_proto_rawDesc), len(file_proto_scorer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_scorer_proto_goTypes,
		DependencyIndexes: file_proto_scorer_proto_depIdxs,
		MessageInfos:      file_proto_scorer_proto_msgTypes,
	}.Build()
	File_proto_scorer_proto = out.File
	file_proto_scorer_proto_goTypes = nil
	file_proto_scorer_proto_depIdxs = nil
}
