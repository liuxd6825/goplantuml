package tags

type Tag interface {
	GetTagType() TagType
}

type TagType string

type BaseTag struct {
	TagType TagType `json:"tagType"`
}

const (
	TagTypeData       TagType = "@data"
	TagTypeAggregate  TagType = "@aggregate"
	TagTypeEnum       TagType = "@enum"
	TagTypeForm       TagType = "@form"
	TagTypeGrid       TagType = "@grid"
	TagTypeRepository TagType = "@repository"
	TagTypeService    TagType = "@service"
	TagTypeEntity     TagType = "@entity"
	TagTypeEvent      TagType = "@event"
	TagTypeApi        TagType = "@api"
	TagTypeView       TagType = "@view"
	TagTypeQuery      TagType = "@query"
	TagTypeCmd        TagType = "@cmd"
)

func (t *BaseTag) GetTagType() TagType {
	return t.TagType
}
