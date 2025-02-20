package appdefinition

import (
	"encoding/json"

	"github.com/acorn-io/aml/cli/pkg/flagargs"
	"github.com/acorn-io/aml/pkg/schema"
	v1 "github.com/acorn-io/runtime/pkg/apis/internal.acorn.io/v1"
)

var (
	hiddenProfiles = map[string]struct{}{
		"devMode":     {},
		"autoUpgrade": {},
	}
	hiddenArgs = map[string]struct{}{
		"dev":         {},
		"autoUpgrade": {},
	}
)

func fromNames(names schema.Names) (result []v1.Profile) {
	for _, name := range names {
		result = append(result, v1.Profile(name))
	}
	return
}

func dropHiddenProfiles(names schema.Names) (result schema.Names) {
	for _, profile := range names {
		if _, ok := hiddenProfiles[profile.Name]; ok {
			continue
		}
		result = append(result, profile)
	}
	return
}

func dropHiddenArgs(args []schema.Field) (result []schema.Field) {
	for _, arg := range args {
		if _, ok := hiddenArgs[arg.Name]; ok {
			continue
		}
		result = append(result, arg)
	}
	return
}

func fromFields(in []schema.Field) (result []v1.Field) {
	for _, item := range in {
		result = append(result, fromField(item))
	}
	return
}

func anyToString(v any) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	d, _ := json.Marshal(v)
	return string(d)
}

func fromObject(in *schema.Object) *v1.Object {
	if in == nil {
		return nil
	}
	return &v1.Object{
		Path:         in.Path,
		Reference:    in.Reference,
		Description:  in.Description,
		Fields:       fromFields(in.Fields),
		AllowNewKeys: in.AllowNewKeys,
	}
}

func fromArray(in *schema.Array) *v1.Array {
	if in == nil {
		return nil
	}
	return &v1.Array{
		Types: fromFieldTypes(in.Types),
	}
}

func fromConstraints(in []schema.Constraint) (result []v1.Constraint) {
	for _, item := range in {
		result = append(result, v1.Constraint{
			Description: item.Description,
			Op:          item.Op,
			Right:       anyToString(item.Right),
			Type:        anyToFieldType(item.Right),
		})
	}
	return
}

func anyToFieldType(v any) *v1.FieldType {
	rt, _ := v.(*v1.FieldType)
	return rt
}

func fromFieldTypes(in []schema.FieldType) (out []v1.FieldType) {
	for _, fieldType := range in {
		out = append(out, *fromFieldType(&fieldType))
	}
	return
}

func fromFieldType(in *schema.FieldType) *v1.FieldType {
	if in == nil {
		return nil
	}
	return &v1.FieldType{
		Kind:        string(in.Kind),
		Object:      fromObject(in.Object),
		Array:       fromArray(in.Array),
		Constraints: fromConstraints(in.Contstraints),
		Default:     anyToString(in.Default),
		Alternates:  fromFieldTypes(in.Alternates),
	}
}

func fromField(in schema.Field) v1.Field {
	return v1.Field{
		Name:        in.Name,
		Description: in.Description,
		Type:        *fromFieldType(&in.Type),
		Match:       in.Match,
		Optional:    in.Optional,
	}
}

type Flags interface {
	Parse(args []string) (map[string]any, []string, error)
}

func (a *AppDefinition) ToFlags(programName, argsFile string, usage func()) (Flags, error) {
	var file schema.File
	err := a.decode(&file)
	if err != nil {
		return nil, err
	}

	args := flagargs.New(argsFile, programName,
		dropHiddenProfiles(file.ProfileNames),
		dropHiddenArgs(file.Args.Fields))
	args.Usage = usage
	return args, nil
}

func (a *AppDefinition) ToParamSpec() (*v1.ParamSpec, error) {
	var file schema.File
	err := a.decode(&file)
	if err != nil {
		return nil, err
	}
	result := &v1.ParamSpec{
		Args:     fromFields(dropHiddenArgs(file.Args.Fields)),
		Profiles: fromNames(dropHiddenProfiles(file.ProfileNames)),
	}

	return result, nil
}
