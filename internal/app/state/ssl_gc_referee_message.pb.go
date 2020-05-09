// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_gc_referee_message.proto

package state

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// These are the "coarse" stages of the game.
type Referee_Stage int32

const (
	// The first half is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	Referee_NORMAL_FIRST_HALF_PRE Referee_Stage = 0
	// The first half of the normal game, before half time.
	Referee_NORMAL_FIRST_HALF Referee_Stage = 1
	// Half time between first and second halves.
	Referee_NORMAL_HALF_TIME Referee_Stage = 2
	// The second half is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	Referee_NORMAL_SECOND_HALF_PRE Referee_Stage = 3
	// The second half of the normal game, after half time.
	Referee_NORMAL_SECOND_HALF Referee_Stage = 4
	// The break before extra time.
	Referee_EXTRA_TIME_BREAK Referee_Stage = 5
	// The first half of extra time is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	Referee_EXTRA_FIRST_HALF_PRE Referee_Stage = 6
	// The first half of extra time.
	Referee_EXTRA_FIRST_HALF Referee_Stage = 7
	// Half time between first and second extra halves.
	Referee_EXTRA_HALF_TIME Referee_Stage = 8
	// The second half of extra time is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	Referee_EXTRA_SECOND_HALF_PRE Referee_Stage = 9
	// The second half of extra time.
	Referee_EXTRA_SECOND_HALF Referee_Stage = 10
	// The break before penalty shootout.
	Referee_PENALTY_SHOOTOUT_BREAK Referee_Stage = 11
	// The penalty shootout.
	Referee_PENALTY_SHOOTOUT Referee_Stage = 12
	// The game is over.
	Referee_POST_GAME Referee_Stage = 13
)

var Referee_Stage_name = map[int32]string{
	0:  "NORMAL_FIRST_HALF_PRE",
	1:  "NORMAL_FIRST_HALF",
	2:  "NORMAL_HALF_TIME",
	3:  "NORMAL_SECOND_HALF_PRE",
	4:  "NORMAL_SECOND_HALF",
	5:  "EXTRA_TIME_BREAK",
	6:  "EXTRA_FIRST_HALF_PRE",
	7:  "EXTRA_FIRST_HALF",
	8:  "EXTRA_HALF_TIME",
	9:  "EXTRA_SECOND_HALF_PRE",
	10: "EXTRA_SECOND_HALF",
	11: "PENALTY_SHOOTOUT_BREAK",
	12: "PENALTY_SHOOTOUT",
	13: "POST_GAME",
}

var Referee_Stage_value = map[string]int32{
	"NORMAL_FIRST_HALF_PRE":  0,
	"NORMAL_FIRST_HALF":      1,
	"NORMAL_HALF_TIME":       2,
	"NORMAL_SECOND_HALF_PRE": 3,
	"NORMAL_SECOND_HALF":     4,
	"EXTRA_TIME_BREAK":       5,
	"EXTRA_FIRST_HALF_PRE":   6,
	"EXTRA_FIRST_HALF":       7,
	"EXTRA_HALF_TIME":        8,
	"EXTRA_SECOND_HALF_PRE":  9,
	"EXTRA_SECOND_HALF":      10,
	"PENALTY_SHOOTOUT_BREAK": 11,
	"PENALTY_SHOOTOUT":       12,
	"POST_GAME":              13,
}

func (x Referee_Stage) Enum() *Referee_Stage {
	p := new(Referee_Stage)
	*p = x
	return p
}

func (x Referee_Stage) String() string {
	return proto.EnumName(Referee_Stage_name, int32(x))
}

func (x *Referee_Stage) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Referee_Stage_value, data, "Referee_Stage")
	if err != nil {
		return err
	}
	*x = Referee_Stage(value)
	return nil
}

func (Referee_Stage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_92a5ea5b58ff1871, []int{0, 0}
}

// These are the "fine" states of play on the field.
type Referee_Command int32

const (
	// All robots should completely stop moving.
	Referee_HALT Referee_Command = 0
	// Robots must keep 50 cm from the ball.
	Referee_STOP Referee_Command = 1
	// A prepared kickoff or penalty may now be taken.
	Referee_NORMAL_START Referee_Command = 2
	// The ball is dropped and free for either team.
	Referee_FORCE_START Referee_Command = 3
	// The yellow team may move into kickoff position.
	Referee_PREPARE_KICKOFF_YELLOW Referee_Command = 4
	// The blue team may move into kickoff position.
	Referee_PREPARE_KICKOFF_BLUE Referee_Command = 5
	// The yellow team may move into penalty position.
	Referee_PREPARE_PENALTY_YELLOW Referee_Command = 6
	// The blue team may move into penalty position.
	Referee_PREPARE_PENALTY_BLUE Referee_Command = 7
	// The yellow team may take a direct free kick.
	Referee_DIRECT_FREE_YELLOW Referee_Command = 8
	// The blue team may take a direct free kick.
	Referee_DIRECT_FREE_BLUE Referee_Command = 9
	// The yellow team may take an indirect free kick.
	Referee_INDIRECT_FREE_YELLOW Referee_Command = 10
	// The blue team may take an indirect free kick.
	Referee_INDIRECT_FREE_BLUE Referee_Command = 11
	// The yellow team is currently in a timeout.
	Referee_TIMEOUT_YELLOW Referee_Command = 12
	// The blue team is currently in a timeout.
	Referee_TIMEOUT_BLUE Referee_Command = 13
	// The yellow team just scored a goal.
	// For information only.
	// For rules compliance, teams must treat as STOP.
	// Deprecated: Use the score field from the team infos instead. That way, you can also detect revoked goals.
	Referee_GOAL_YELLOW Referee_Command = 14 // Deprecated: Do not use.
	// The blue team just scored a goal. See also GOAL_YELLOW.
	Referee_GOAL_BLUE Referee_Command = 15 // Deprecated: Do not use.
	// Equivalent to STOP, but the yellow team must pick up the ball and
	// drop it in the Designated Position.
	Referee_BALL_PLACEMENT_YELLOW Referee_Command = 16
	// Equivalent to STOP, but the blue team must pick up the ball and drop
	// it in the Designated Position.
	Referee_BALL_PLACEMENT_BLUE Referee_Command = 17
)

var Referee_Command_name = map[int32]string{
	0:  "HALT",
	1:  "STOP",
	2:  "NORMAL_START",
	3:  "FORCE_START",
	4:  "PREPARE_KICKOFF_YELLOW",
	5:  "PREPARE_KICKOFF_BLUE",
	6:  "PREPARE_PENALTY_YELLOW",
	7:  "PREPARE_PENALTY_BLUE",
	8:  "DIRECT_FREE_YELLOW",
	9:  "DIRECT_FREE_BLUE",
	10: "INDIRECT_FREE_YELLOW",
	11: "INDIRECT_FREE_BLUE",
	12: "TIMEOUT_YELLOW",
	13: "TIMEOUT_BLUE",
	14: "GOAL_YELLOW",
	15: "GOAL_BLUE",
	16: "BALL_PLACEMENT_YELLOW",
	17: "BALL_PLACEMENT_BLUE",
}

var Referee_Command_value = map[string]int32{
	"HALT":                   0,
	"STOP":                   1,
	"NORMAL_START":           2,
	"FORCE_START":            3,
	"PREPARE_KICKOFF_YELLOW": 4,
	"PREPARE_KICKOFF_BLUE":   5,
	"PREPARE_PENALTY_YELLOW": 6,
	"PREPARE_PENALTY_BLUE":   7,
	"DIRECT_FREE_YELLOW":     8,
	"DIRECT_FREE_BLUE":       9,
	"INDIRECT_FREE_YELLOW":   10,
	"INDIRECT_FREE_BLUE":     11,
	"TIMEOUT_YELLOW":         12,
	"TIMEOUT_BLUE":           13,
	"GOAL_YELLOW":            14,
	"GOAL_BLUE":              15,
	"BALL_PLACEMENT_YELLOW":  16,
	"BALL_PLACEMENT_BLUE":    17,
}

func (x Referee_Command) Enum() *Referee_Command {
	p := new(Referee_Command)
	*p = x
	return p
}

func (x Referee_Command) String() string {
	return proto.EnumName(Referee_Command_name, int32(x))
}

func (x *Referee_Command) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Referee_Command_value, data, "Referee_Command")
	if err != nil {
		return err
	}
	*x = Referee_Command(value)
	return nil
}

func (Referee_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_92a5ea5b58ff1871, []int{0, 1}
}

// Each UDP packet contains one of these messages.
type Referee struct {
	// The UNIX timestamp when the packet was sent, in microseconds.
	// Divide by 1,000,000 to get a time_t.
	PacketTimestamp *uint64        `protobuf:"varint,1,req,name=packet_timestamp,json=packetTimestamp" json:"packet_timestamp,omitempty"`
	Stage           *Referee_Stage `protobuf:"varint,2,req,name=stage,enum=Referee_Stage" json:"stage,omitempty"`
	// The number of microseconds left in the stage.
	// The following stages have this value; the rest do not:
	// NORMAL_FIRST_HALF
	// NORMAL_HALF_TIME
	// NORMAL_SECOND_HALF
	// EXTRA_TIME_BREAK
	// EXTRA_FIRST_HALF
	// EXTRA_HALF_TIME
	// EXTRA_SECOND_HALF
	// PENALTY_SHOOTOUT_BREAK
	//
	// If the stage runs over its specified time, this value
	// becomes negative.
	StageTimeLeft *int32           `protobuf:"zigzag32,3,opt,name=stage_time_left,json=stageTimeLeft" json:"stage_time_left,omitempty"`
	Command       *Referee_Command `protobuf:"varint,4,req,name=command,enum=Referee_Command" json:"command,omitempty"`
	// The number of commands issued since startup (mod 2^32).
	CommandCounter *uint32 `protobuf:"varint,5,req,name=command_counter,json=commandCounter" json:"command_counter,omitempty"`
	// The UNIX timestamp when the command was issued, in microseconds.
	// This value changes only when a new command is issued, not on each packet.
	CommandTimestamp *uint64 `protobuf:"varint,6,req,name=command_timestamp,json=commandTimestamp" json:"command_timestamp,omitempty"`
	// Information about the two teams.
	Yellow             *Referee_TeamInfo `protobuf:"bytes,7,req,name=yellow" json:"yellow,omitempty"`
	Blue               *Referee_TeamInfo `protobuf:"bytes,8,req,name=blue" json:"blue,omitempty"`
	DesignatedPosition *Referee_Point    `protobuf:"bytes,9,opt,name=designated_position,json=designatedPosition" json:"designated_position,omitempty"`
	// Information about the direction of play.
	// True, if the blue team will have it's goal on the positive x-axis of the ssl-vision coordinate system.
	// Obviously, the yellow team will play on the opposite half.
	BlueTeamOnPositiveHalf *bool `protobuf:"varint,10,opt,name=blue_team_on_positive_half,json=blueTeamOnPositiveHalf" json:"blue_team_on_positive_half,omitempty"`
	// The command that will be issued after the current stoppage and ball placement to continue the game.
	NextCommand        *Referee_Command          `protobuf:"varint,12,opt,name=next_command,json=nextCommand,enum=Referee_Command" json:"next_command,omitempty"`
	GameEvents         []*GameEvent              `protobuf:"bytes,16,rep,name=game_events,json=gameEvents" json:"game_events,omitempty"`
	GameEventProposals []*GameEventProposalGroup `protobuf:"bytes,17,rep,name=game_event_proposals,json=gameEventProposals" json:"game_event_proposals,omitempty"`
	// The time in microseconds that is remaining until the current action times out
	// The time will not be reset. It can get negative.
	// An autoRef would raise an appropriate event, if the time gets negative.
	// Possible actions where this time is relevant:
	//  * free kicks
	//  * kickoff, penalty kick, force start
	//  * ball placement
	CurrentActionTimeRemaining *int32   `protobuf:"varint,15,opt,name=current_action_time_remaining,json=currentActionTimeRemaining" json:"current_action_time_remaining,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Referee) Reset()         { *m = Referee{} }
func (m *Referee) String() string { return proto.CompactTextString(m) }
func (*Referee) ProtoMessage()    {}
func (*Referee) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a5ea5b58ff1871, []int{0}
}

func (m *Referee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Referee.Unmarshal(m, b)
}
func (m *Referee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Referee.Marshal(b, m, deterministic)
}
func (m *Referee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Referee.Merge(m, src)
}
func (m *Referee) XXX_Size() int {
	return xxx_messageInfo_Referee.Size(m)
}
func (m *Referee) XXX_DiscardUnknown() {
	xxx_messageInfo_Referee.DiscardUnknown(m)
}

var xxx_messageInfo_Referee proto.InternalMessageInfo

func (m *Referee) GetPacketTimestamp() uint64 {
	if m != nil && m.PacketTimestamp != nil {
		return *m.PacketTimestamp
	}
	return 0
}

func (m *Referee) GetStage() Referee_Stage {
	if m != nil && m.Stage != nil {
		return *m.Stage
	}
	return Referee_NORMAL_FIRST_HALF_PRE
}

func (m *Referee) GetStageTimeLeft() int32 {
	if m != nil && m.StageTimeLeft != nil {
		return *m.StageTimeLeft
	}
	return 0
}

func (m *Referee) GetCommand() Referee_Command {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return Referee_HALT
}

func (m *Referee) GetCommandCounter() uint32 {
	if m != nil && m.CommandCounter != nil {
		return *m.CommandCounter
	}
	return 0
}

func (m *Referee) GetCommandTimestamp() uint64 {
	if m != nil && m.CommandTimestamp != nil {
		return *m.CommandTimestamp
	}
	return 0
}

func (m *Referee) GetYellow() *Referee_TeamInfo {
	if m != nil {
		return m.Yellow
	}
	return nil
}

func (m *Referee) GetBlue() *Referee_TeamInfo {
	if m != nil {
		return m.Blue
	}
	return nil
}

func (m *Referee) GetDesignatedPosition() *Referee_Point {
	if m != nil {
		return m.DesignatedPosition
	}
	return nil
}

func (m *Referee) GetBlueTeamOnPositiveHalf() bool {
	if m != nil && m.BlueTeamOnPositiveHalf != nil {
		return *m.BlueTeamOnPositiveHalf
	}
	return false
}

func (m *Referee) GetNextCommand() Referee_Command {
	if m != nil && m.NextCommand != nil {
		return *m.NextCommand
	}
	return Referee_HALT
}

func (m *Referee) GetGameEvents() []*GameEvent {
	if m != nil {
		return m.GameEvents
	}
	return nil
}

func (m *Referee) GetGameEventProposals() []*GameEventProposalGroup {
	if m != nil {
		return m.GameEventProposals
	}
	return nil
}

func (m *Referee) GetCurrentActionTimeRemaining() int32 {
	if m != nil && m.CurrentActionTimeRemaining != nil {
		return *m.CurrentActionTimeRemaining
	}
	return 0
}

// Information about a single team.
type Referee_TeamInfo struct {
	// The team's name (empty string if operator has not typed anything).
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// The number of goals scored by the team during normal play and overtime.
	Score *uint32 `protobuf:"varint,2,req,name=score" json:"score,omitempty"`
	// The number of red cards issued to the team since the beginning of the game.
	RedCards *uint32 `protobuf:"varint,3,req,name=red_cards,json=redCards" json:"red_cards,omitempty"`
	// The amount of time (in microseconds) left on each yellow card issued to the team.
	// If no yellow cards are issued, this array has no elements.
	// Otherwise, times are ordered from smallest to largest.
	YellowCardTimes []uint32 `protobuf:"varint,4,rep,packed,name=yellow_card_times,json=yellowCardTimes" json:"yellow_card_times,omitempty"`
	// The total number of yellow cards ever issued to the team.
	YellowCards *uint32 `protobuf:"varint,5,req,name=yellow_cards,json=yellowCards" json:"yellow_cards,omitempty"`
	// The number of timeouts this team can still call.
	// If in a timeout right now, that timeout is excluded.
	Timeouts *uint32 `protobuf:"varint,6,req,name=timeouts" json:"timeouts,omitempty"`
	// The number of microseconds of timeout this team can use.
	TimeoutTime *uint32 `protobuf:"varint,7,req,name=timeout_time,json=timeoutTime" json:"timeout_time,omitempty"`
	// The pattern number of this team's goalkeeper.
	Goalkeeper *uint32 `protobuf:"varint,8,req,name=goalkeeper" json:"goalkeeper,omitempty"`
	// The total number of countable fouls that act towards yellow cards
	FoulCounter *uint32 `protobuf:"varint,9,opt,name=foul_counter,json=foulCounter" json:"foul_counter,omitempty"`
	// The number of consecutive ball placement failures of this team
	BallPlacementFailures *uint32 `protobuf:"varint,10,opt,name=ball_placement_failures,json=ballPlacementFailures" json:"ball_placement_failures,omitempty"`
	// Indicate if the team is able and allowed to place the ball
	CanPlaceBall *bool `protobuf:"varint,12,opt,name=can_place_ball,json=canPlaceBall" json:"can_place_ball,omitempty"`
	// The maximum number of bots allowed on the field based on division and cards
	MaxAllowedBots *uint32 `protobuf:"varint,13,opt,name=max_allowed_bots,json=maxAllowedBots" json:"max_allowed_bots,omitempty"`
	// The team has submitted an intent to substitute one or more robots at the next chance
	BotSubstitutionIntent *bool `protobuf:"varint,14,opt,name=bot_substitution_intent,json=botSubstitutionIntent" json:"bot_substitution_intent,omitempty"`
	// Indicate if the team reached the maximum allowed ball placement failures and is thus not allowed to place the ball anymore
	BallPlacementFailuresReached *bool    `protobuf:"varint,15,opt,name=ball_placement_failures_reached,json=ballPlacementFailuresReached" json:"ball_placement_failures_reached,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *Referee_TeamInfo) Reset()         { *m = Referee_TeamInfo{} }
func (m *Referee_TeamInfo) String() string { return proto.CompactTextString(m) }
func (*Referee_TeamInfo) ProtoMessage()    {}
func (*Referee_TeamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a5ea5b58ff1871, []int{0, 0}
}

func (m *Referee_TeamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Referee_TeamInfo.Unmarshal(m, b)
}
func (m *Referee_TeamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Referee_TeamInfo.Marshal(b, m, deterministic)
}
func (m *Referee_TeamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Referee_TeamInfo.Merge(m, src)
}
func (m *Referee_TeamInfo) XXX_Size() int {
	return xxx_messageInfo_Referee_TeamInfo.Size(m)
}
func (m *Referee_TeamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Referee_TeamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Referee_TeamInfo proto.InternalMessageInfo

func (m *Referee_TeamInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Referee_TeamInfo) GetScore() uint32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *Referee_TeamInfo) GetRedCards() uint32 {
	if m != nil && m.RedCards != nil {
		return *m.RedCards
	}
	return 0
}

func (m *Referee_TeamInfo) GetYellowCardTimes() []uint32 {
	if m != nil {
		return m.YellowCardTimes
	}
	return nil
}

func (m *Referee_TeamInfo) GetYellowCards() uint32 {
	if m != nil && m.YellowCards != nil {
		return *m.YellowCards
	}
	return 0
}

func (m *Referee_TeamInfo) GetTimeouts() uint32 {
	if m != nil && m.Timeouts != nil {
		return *m.Timeouts
	}
	return 0
}

func (m *Referee_TeamInfo) GetTimeoutTime() uint32 {
	if m != nil && m.TimeoutTime != nil {
		return *m.TimeoutTime
	}
	return 0
}

func (m *Referee_TeamInfo) GetGoalkeeper() uint32 {
	if m != nil && m.Goalkeeper != nil {
		return *m.Goalkeeper
	}
	return 0
}

func (m *Referee_TeamInfo) GetFoulCounter() uint32 {
	if m != nil && m.FoulCounter != nil {
		return *m.FoulCounter
	}
	return 0
}

func (m *Referee_TeamInfo) GetBallPlacementFailures() uint32 {
	if m != nil && m.BallPlacementFailures != nil {
		return *m.BallPlacementFailures
	}
	return 0
}

func (m *Referee_TeamInfo) GetCanPlaceBall() bool {
	if m != nil && m.CanPlaceBall != nil {
		return *m.CanPlaceBall
	}
	return false
}

func (m *Referee_TeamInfo) GetMaxAllowedBots() uint32 {
	if m != nil && m.MaxAllowedBots != nil {
		return *m.MaxAllowedBots
	}
	return 0
}

func (m *Referee_TeamInfo) GetBotSubstitutionIntent() bool {
	if m != nil && m.BotSubstitutionIntent != nil {
		return *m.BotSubstitutionIntent
	}
	return false
}

func (m *Referee_TeamInfo) GetBallPlacementFailuresReached() bool {
	if m != nil && m.BallPlacementFailuresReached != nil {
		return *m.BallPlacementFailuresReached
	}
	return false
}

// The coordinates of the Designated Position. These are measured in
// millimetres and correspond to SSL-Vision coordinates. These fields are
// always either both present (in the case of a ball placement command) or
// both absent (in the case of any other command).
type Referee_Point struct {
	X                    *float32 `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y                    *float32 `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Referee_Point) Reset()         { *m = Referee_Point{} }
func (m *Referee_Point) String() string { return proto.CompactTextString(m) }
func (*Referee_Point) ProtoMessage()    {}
func (*Referee_Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a5ea5b58ff1871, []int{0, 1}
}

func (m *Referee_Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Referee_Point.Unmarshal(m, b)
}
func (m *Referee_Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Referee_Point.Marshal(b, m, deterministic)
}
func (m *Referee_Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Referee_Point.Merge(m, src)
}
func (m *Referee_Point) XXX_Size() int {
	return xxx_messageInfo_Referee_Point.Size(m)
}
func (m *Referee_Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Referee_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Referee_Point proto.InternalMessageInfo

func (m *Referee_Point) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Referee_Point) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

// List of matching proposals
type GameEventProposalGroup struct {
	// The proposed game event.
	GameEvent []*GameEvent `protobuf:"bytes,1,rep,name=game_event,json=gameEvent" json:"game_event,omitempty"`
	// Whether the proposal group was accepted
	Accepted             *bool    `protobuf:"varint,2,opt,name=accepted" json:"accepted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameEventProposalGroup) Reset()         { *m = GameEventProposalGroup{} }
func (m *GameEventProposalGroup) String() string { return proto.CompactTextString(m) }
func (*GameEventProposalGroup) ProtoMessage()    {}
func (*GameEventProposalGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a5ea5b58ff1871, []int{1}
}

func (m *GameEventProposalGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameEventProposalGroup.Unmarshal(m, b)
}
func (m *GameEventProposalGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameEventProposalGroup.Marshal(b, m, deterministic)
}
func (m *GameEventProposalGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameEventProposalGroup.Merge(m, src)
}
func (m *GameEventProposalGroup) XXX_Size() int {
	return xxx_messageInfo_GameEventProposalGroup.Size(m)
}
func (m *GameEventProposalGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_GameEventProposalGroup.DiscardUnknown(m)
}

var xxx_messageInfo_GameEventProposalGroup proto.InternalMessageInfo

func (m *GameEventProposalGroup) GetGameEvent() []*GameEvent {
	if m != nil {
		return m.GameEvent
	}
	return nil
}

func (m *GameEventProposalGroup) GetAccepted() bool {
	if m != nil && m.Accepted != nil {
		return *m.Accepted
	}
	return false
}

func init() {
	proto.RegisterEnum("Referee_Stage", Referee_Stage_name, Referee_Stage_value)
	proto.RegisterEnum("Referee_Command", Referee_Command_name, Referee_Command_value)
	proto.RegisterType((*Referee)(nil), "Referee")
	proto.RegisterType((*Referee_TeamInfo)(nil), "Referee.TeamInfo")
	proto.RegisterType((*Referee_Point)(nil), "Referee.Point")
	proto.RegisterType((*GameEventProposalGroup)(nil), "GameEventProposalGroup")
}

func init() {
	proto.RegisterFile("ssl_gc_referee_message.proto", fileDescriptor_92a5ea5b58ff1871)
}

var fileDescriptor_92a5ea5b58ff1871 = []byte{
	// 1174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xed, 0x6e, 0xdb, 0x36,
	0x14, 0xad, 0xfc, 0x91, 0xc8, 0xd7, 0x96, 0x2d, 0x33, 0x5f, 0xaa, 0xd7, 0x6d, 0x5e, 0xd6, 0x6d,
	0xee, 0x8a, 0x26, 0x40, 0x07, 0xec, 0xc7, 0x80, 0xa1, 0x90, 0x5d, 0x39, 0x71, 0xeb, 0xc4, 0x06,
	0xad, 0x62, 0xeb, 0xfe, 0x10, 0xb4, 0x4c, 0xbb, 0x46, 0x25, 0xd1, 0x10, 0xe9, 0x2e, 0x79, 0x89,
	0xed, 0x49, 0xf6, 0x30, 0x7b, 0xa3, 0x81, 0x94, 0x64, 0x7b, 0x49, 0xfa, 0x8f, 0x3c, 0xe7, 0xdc,
	0xa3, 0x7b, 0x2f, 0x2f, 0x29, 0x78, 0x22, 0x44, 0x48, 0x16, 0x01, 0x49, 0xd8, 0x9c, 0x25, 0x8c,
	0x91, 0x88, 0x09, 0x41, 0x17, 0xec, 0x6c, 0x95, 0x70, 0xc9, 0x5b, 0x27, 0x19, 0xbb, 0xa0, 0x11,
	0x23, 0xec, 0x13, 0x8b, 0x65, 0x4a, 0x9c, 0xfe, 0x6d, 0xc3, 0x3e, 0x4e, 0x43, 0xd0, 0x33, 0xb0,
	0x57, 0x34, 0xf8, 0xc8, 0x24, 0x91, 0xcb, 0x88, 0x09, 0x49, 0xa3, 0x95, 0x63, 0xb4, 0x0b, 0x9d,
	0x12, 0x6e, 0xa4, 0xb8, 0x9f, 0xc3, 0xe8, 0x29, 0x94, 0x85, 0xa4, 0x0b, 0xe6, 0x14, 0xda, 0x85,
	0x4e, 0xfd, 0x65, 0xfd, 0x2c, 0xf3, 0x38, 0x9b, 0x28, 0x14, 0xa7, 0x24, 0xfa, 0x1e, 0x1a, 0x7a,
	0xa1, 0xfd, 0x48, 0xc8, 0xe6, 0xd2, 0x29, 0xb6, 0x8d, 0x4e, 0x13, 0x5b, 0x1a, 0x56, 0x76, 0x43,
	0x36, 0x97, 0xe8, 0x47, 0xd8, 0x0f, 0x78, 0x14, 0xd1, 0x78, 0xe6, 0x94, 0xb4, 0x9f, 0xbd, 0xf1,
	0xeb, 0xa5, 0x38, 0xce, 0x05, 0xe8, 0x07, 0x68, 0x64, 0x4b, 0x12, 0xf0, 0x75, 0x2c, 0x59, 0xe2,
	0x94, 0xdb, 0x85, 0x8e, 0x85, 0xeb, 0x19, 0xdc, 0x4b, 0x51, 0xf4, 0x1c, 0x9a, 0xb9, 0x70, 0x5b,
	0xce, 0x9e, 0x2e, 0xc7, 0xce, 0x88, 0x6d, 0x3d, 0xcf, 0x60, 0xef, 0x96, 0x85, 0x21, 0xff, 0xd3,
	0xd9, 0x6f, 0x17, 0x3a, 0xd5, 0x97, 0xcd, 0x4d, 0x02, 0x3e, 0xa3, 0xd1, 0x20, 0x9e, 0x73, 0x9c,
	0x09, 0xd0, 0x77, 0x50, 0x9a, 0x86, 0x6b, 0xe6, 0x98, 0x9f, 0x13, 0x6a, 0x1a, 0xbd, 0x82, 0x83,
	0x19, 0x13, 0xcb, 0x45, 0x4c, 0x25, 0x9b, 0x91, 0x15, 0x17, 0x4b, 0xb9, 0xe4, 0xb1, 0x53, 0x69,
	0x1b, 0x9d, 0xea, 0x4e, 0xbf, 0xc6, 0x7c, 0x19, 0x4b, 0x8c, 0xb6, 0xd2, 0x71, 0xa6, 0x44, 0xbf,
	0x40, 0x4b, 0x19, 0x11, 0xc9, 0x68, 0x44, 0x78, 0x9c, 0x59, 0x7c, 0x62, 0xe4, 0x03, 0x0d, 0xe7,
	0x0e, 0xb4, 0x8d, 0x8e, 0x89, 0x8f, 0x95, 0x42, 0x7d, 0x78, 0x14, 0x8f, 0x33, 0xfa, 0x92, 0x86,
	0x73, 0xf4, 0x13, 0xd4, 0x62, 0x76, 0x23, 0x49, 0xde, 0xd5, 0x5a, 0xdb, 0x78, 0xb0, 0xab, 0x55,
	0xa5, 0xca, 0x36, 0xe8, 0x39, 0x54, 0xb7, 0xe3, 0x21, 0x1c, 0xbb, 0x5d, 0xec, 0x54, 0x5f, 0xc2,
	0xd9, 0x05, 0x8d, 0x98, 0xa7, 0x20, 0x0c, 0x8b, 0x7c, 0x29, 0xd0, 0x00, 0x0e, 0xb7, 0x62, 0xb2,
	0x4a, 0xf8, 0x8a, 0x0b, 0x1a, 0x0a, 0xa7, 0xa9, 0xa3, 0x4e, 0xb6, 0x51, 0xe3, 0x8c, 0xba, 0x48,
	0xf8, 0x7a, 0x85, 0xd1, 0xe2, 0x2e, 0x2e, 0x90, 0x0b, 0x5f, 0x06, 0xeb, 0x24, 0x51, 0x3e, 0x34,
	0x50, 0xa5, 0xa7, 0xe3, 0x92, 0xb0, 0x88, 0x2e, 0xe3, 0x65, 0xbc, 0x70, 0x1a, 0x6d, 0xa3, 0x53,
	0xc6, 0xad, 0x4c, 0xe4, 0x6a, 0x8d, 0x3a, 0x3a, 0x9c, 0x2b, 0x5a, 0xff, 0x94, 0xc0, 0xcc, 0xfb,
	0x8f, 0x10, 0x94, 0x62, 0x1a, 0x31, 0x3d, 0xba, 0x15, 0xac, 0xd7, 0xe8, 0x10, 0xca, 0x22, 0xe0,
	0x49, 0x3a, 0xaf, 0x16, 0x4e, 0x37, 0xe8, 0x0b, 0xa8, 0x24, 0x6c, 0x46, 0x02, 0x9a, 0xcc, 0x84,
	0x53, 0xd4, 0x8c, 0x99, 0xb0, 0x59, 0x4f, 0xed, 0xd1, 0x19, 0x34, 0xd3, 0x13, 0xd7, 0x7c, 0x3a,
	0x43, 0x4e, 0xa9, 0x5d, 0xec, 0x58, 0xdd, 0x82, 0x6d, 0xe0, 0x46, 0x4a, 0x2a, 0xad, 0x1e, 0x23,
	0xf4, 0x0d, 0xd4, 0x76, 0xf4, 0x22, 0x9b, 0xca, 0xea, 0x56, 0x26, 0x50, 0x0b, 0x4c, 0x65, 0xc3,
	0xd7, 0x52, 0xe8, 0x49, 0xb4, 0xf0, 0x66, 0xaf, 0xc2, 0xb3, 0xb5, 0xfe, 0x94, 0x9e, 0x43, 0x0b,
	0x57, 0x33, 0x4c, 0x7d, 0x02, 0x7d, 0x05, 0xb0, 0xe0, 0x34, 0xfc, 0xc8, 0xd8, 0x8a, 0x25, 0x7a,
	0xfe, 0x2c, 0xbc, 0x83, 0x28, 0x8b, 0x39, 0x5f, 0x87, 0x9b, 0x7b, 0xa1, 0x66, 0xcd, 0xc2, 0x55,
	0x85, 0xe5, 0x97, 0xe2, 0x67, 0x38, 0x99, 0xd2, 0x30, 0x24, 0xab, 0x90, 0x06, 0x2c, 0x52, 0x2d,
	0x9f, 0xd3, 0x65, 0xb8, 0x4e, 0x98, 0xd0, 0x13, 0x65, 0xe1, 0x23, 0x45, 0x8f, 0x73, 0xb6, 0x9f,
	0x91, 0xe8, 0x29, 0xd4, 0x03, 0x1a, 0xa7, 0x61, 0x44, 0x49, 0xf4, 0x48, 0x99, 0xb8, 0x16, 0xd0,
	0x58, 0xab, 0xbb, 0x34, 0x0c, 0x51, 0x07, 0xec, 0x88, 0xde, 0x10, 0xaa, 0x2a, 0x66, 0x33, 0x32,
	0xe5, 0x52, 0x38, 0x96, 0xb6, 0xad, 0x47, 0xf4, 0xc6, 0x4d, 0xe1, 0x2e, 0x97, 0x42, 0xe7, 0xc1,
	0x25, 0x11, 0xeb, 0xa9, 0x90, 0x4b, 0xb9, 0xd6, 0xa7, 0xbe, 0x8c, 0x25, 0x8b, 0xa5, 0x53, 0xd7,
	0xc6, 0x47, 0x53, 0x2e, 0x27, 0x3b, 0xec, 0x40, 0x93, 0xc8, 0x83, 0xaf, 0x3f, 0x93, 0x3f, 0x49,
	0x18, 0x0d, 0x3e, 0xb0, 0x99, 0x9e, 0x16, 0x13, 0x3f, 0x79, 0xb0, 0x0e, 0x9c, 0x6a, 0x5a, 0xdf,
	0x42, 0x59, 0x5f, 0x3c, 0x54, 0x03, 0xe3, 0x46, 0x0f, 0x4a, 0x01, 0x1b, 0x37, 0x6a, 0x77, 0xab,
	0x27, 0xa4, 0x80, 0x8d, 0xdb, 0xd3, 0x7f, 0x0b, 0x50, 0xd6, 0xcf, 0x19, 0x7a, 0x0c, 0x47, 0xd7,
	0x23, 0x7c, 0xe5, 0x0e, 0x49, 0x7f, 0x80, 0x27, 0x3e, 0xb9, 0x74, 0x87, 0x7d, 0x32, 0xc6, 0x9e,
	0xfd, 0x08, 0x1d, 0x41, 0xf3, 0x1e, 0x65, 0x1b, 0xe8, 0x10, 0xec, 0x0c, 0xd6, 0x5a, 0x7f, 0x70,
	0xe5, 0xd9, 0x05, 0xd4, 0x82, 0xe3, 0x0c, 0x9d, 0x78, 0xbd, 0xd1, 0xf5, 0xeb, 0xad, 0x51, 0x11,
	0x1d, 0x03, 0xba, 0xcf, 0xd9, 0x25, 0xe5, 0xe4, 0xfd, 0xee, 0x63, 0x57, 0x7b, 0x90, 0x2e, 0xf6,
	0xdc, 0xb7, 0x76, 0x19, 0x39, 0x70, 0x98, 0xa2, 0x77, 0x12, 0xda, 0xdb, 0xea, 0x77, 0xf2, 0xd9,
	0x47, 0x07, 0xd0, 0x48, 0xd1, 0x6d, 0x3a, 0xa6, 0x2a, 0x2b, 0x05, 0xef, 0x66, 0x53, 0x51, 0x65,
	0xdd, 0xa3, 0x6c, 0x50, 0x05, 0x8c, 0xbd, 0x6b, 0x77, 0xe8, 0xbf, 0x27, 0x93, 0xcb, 0xd1, 0xc8,
	0x1f, 0xbd, 0xf3, 0xb3, 0x94, 0xaa, 0xea, 0xc3, 0x77, 0x39, 0xbb, 0x86, 0x2c, 0xa8, 0x8c, 0x47,
	0x13, 0x9f, 0x5c, 0xb8, 0x57, 0x9e, 0x6d, 0x9d, 0xfe, 0x55, 0x84, 0xfd, 0xfc, 0xbd, 0x31, 0xa1,
	0x74, 0xe9, 0x0e, 0x7d, 0xfb, 0x91, 0x5a, 0x4d, 0xfc, 0xd1, 0xd8, 0x36, 0x90, 0x0d, 0xb5, 0xbc,
	0x0b, 0xbe, 0x8b, 0x7d, 0xbb, 0x80, 0x1a, 0x50, 0xed, 0x8f, 0x70, 0xcf, 0xcb, 0x80, 0xa2, 0xce,
	0x01, 0x7b, 0x63, 0x17, 0x7b, 0xe4, 0xed, 0xa0, 0xf7, 0x76, 0xd4, 0xef, 0x93, 0xf7, 0xde, 0x70,
	0x38, 0xfa, 0xcd, 0x2e, 0xa9, 0xb6, 0xdc, 0xe5, 0xba, 0xc3, 0x77, 0x9e, 0x5d, 0xde, 0x8d, 0xca,
	0xb3, 0xcc, 0xa2, 0xf6, 0x76, 0xa3, 0x72, 0x4e, 0x47, 0xed, 0xab, 0x43, 0x79, 0x3d, 0xc0, 0x5e,
	0xcf, 0x27, 0x7d, 0xec, 0x79, 0x79, 0x84, 0xa9, 0x6a, 0xdd, 0xc5, 0xb5, 0xba, 0xa2, 0x7c, 0x06,
	0xd7, 0x0f, 0xe8, 0x41, 0xf9, 0xfc, 0x9f, 0xd1, 0x11, 0x55, 0x84, 0xa0, 0xae, 0xce, 0x42, 0xb5,
	0x31, 0xd3, 0xd6, 0x54, 0x0b, 0x72, 0x4c, 0xab, 0x2c, 0x74, 0x00, 0xd5, 0x8b, 0x91, 0x3b, 0xcc,
	0x25, 0xf5, 0x56, 0xc1, 0x34, 0x50, 0x13, 0x2a, 0x1a, 0xd4, 0x9a, 0x86, 0x86, 0x1e, 0xc3, 0x51,
	0xd7, 0x1d, 0x0e, 0xc9, 0x78, 0xe8, 0xf6, 0xbc, 0x2b, 0xef, 0x7a, 0x63, 0x6a, 0xa3, 0x13, 0x38,
	0xb8, 0x43, 0xe9, 0xb8, 0xe6, 0x9b, 0x92, 0x59, 0xb5, 0x6b, 0x6f, 0x4a, 0xa6, 0x65, 0xd7, 0xdf,
	0x94, 0xcc, 0xba, 0xdd, 0x38, 0x25, 0x70, 0xfc, 0xf0, 0xe3, 0x8d, 0x9e, 0x01, 0x6c, 0xdf, 0x7c,
	0xc7, 0xb8, 0xf7, 0x7f, 0xa8, 0x6c, 0x1e, 0x77, 0xf5, 0xd2, 0xd1, 0x20, 0x60, 0x2b, 0xc9, 0x66,
	0x4e, 0x41, 0x5f, 0xc8, 0xcd, 0xbe, 0xfb, 0xea, 0x8f, 0x5f, 0x17, 0x4b, 0xf9, 0x61, 0x3d, 0x3d,
	0x0b, 0x78, 0x74, 0x8e, 0xf9, 0x94, 0xf7, 0xd6, 0xab, 0x17, 0x93, 0xc9, 0xf0, 0x5c, 0x88, 0xf0,
	0x85, 0xf2, 0x78, 0x11, 0xf0, 0x58, 0x26, 0x3c, 0x0c, 0x59, 0x72, 0xae, 0x5e, 0x85, 0x24, 0xa6,
	0xe1, 0x39, 0x5d, 0xad, 0xce, 0x85, 0xa4, 0x92, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xc4,
	0x33, 0x76, 0xeb, 0x08, 0x00, 0x00,
}
