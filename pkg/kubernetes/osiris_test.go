package kubernetes

import (
	"testing"
)

func TestAnnotationBooleanValue(t *testing.T) {
	testcases := []struct {
		name           string
		annotations    map[string]string
		expectedResult bool
	}{
		{
			name: "map with osiris enabled entry and value 1",
			annotations: map[string]string{
				enableScalingAnnotationName: "1",
			},
			expectedResult: true,
		},
		{
			name: "map with osiris enabled entry and value true",
			annotations: map[string]string{
				enableScalingAnnotationName: "true",
			},
			expectedResult: true,
		},
		{
			name: "map with osiris enabled entry and value on",
			annotations: map[string]string{
				enableScalingAnnotationName: "on",
			},
			expectedResult: true,
		},
		{
			name: "map with osiris enabled entry and value y",
			annotations: map[string]string{
				enableScalingAnnotationName: "y",
			},
			expectedResult: true,
		},
		{
			name: "map with osiris enabled entry and value yes",
			annotations: map[string]string{
				enableScalingAnnotationName: "yes",
			},
			expectedResult: true,
		},
		{
			name:           "map with no osiris enabled entry ",
			annotations:    map[string]string{},
			expectedResult: false,
		},

		{
			name: "map with osiris enabled entry and invalid value",
			annotations: map[string]string{
				enableScalingAnnotationName: "yee",
			},
			expectedResult: false,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			actual := annotationBooleanValue(test.annotations, enableScalingAnnotationName)
			if actual != test.expectedResult {
				t.Errorf(
					"expected annotationBooleanValue to return %t, but got %t",
					test.expectedResult, actual)
			}
		})
	}

}

func TestGetMinReplicas(t *testing.T) {
	testcases := []struct {
		name           string
		annotations    map[string]string
		expectedResult int32
	}{
		{
			name: "map with min replicas entry",
			annotations: map[string]string{
				"osiris.dm.gg/minReplicas": "3",
			},
			expectedResult: 3,
		},
		{
			name: "map with no min replicas entry",
			annotations: map[string]string{
				"osiris.dm.gg/notminReplicas": "3",
			},
			expectedResult: 1,
		},
		{
			name: "map with invalid min replicas entry",
			annotations: map[string]string{
				"osiris.dm.gg/minReplicas": "invalid",
			},
			expectedResult: 1,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			actual := GetMinReplicas(test.annotations, 1)
			if actual != test.expectedResult {
				t.Errorf(
					"expected GetMinReplicas to return %d, but got %d",
					test.expectedResult, actual)
			}
		})
	}
}
