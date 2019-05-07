package framework

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	"os"
	"testing"
)

type testCase struct {
	desc          string
	node          *v1.Node
	silent        bool
	conditionType v1.NodeConditionType
	wantTrue      bool
	expect        bool
}

var testCases []testCase

func Test_isNodeConditionSetAsExpectedOriginal(t *testing.T) {
	Logf("Original------------------------")

	for i, tc := range testCases {
		Logf("testcase:[%v]", i)
		result := isNodeConditionSetAsExpected(tc.node, tc.conditionType, tc.wantTrue, tc.silent)
		if tc.expect != result {
			t.Errorf("testcase:[%v] error occurs: %v", i, tc.desc)
		}
	}
}

func Test_isNodeConditionSetAsExpected(t *testing.T) {

	/*testCases2 := []struct {
	  	desc          string
	  	node          *v1.Node
	  	silent        bool
	  	conditionType v1.NodeConditionType
	  	wantTrue      bool
	  	expect        bool
	  }{
	  	{
	  		desc: "Return should be true.  wantTrue: true, conditionType equls node's one, and node has no taints",
	  		node: &v1.Node{
	  			TypeMeta: metav1.TypeMeta{},
	  			ObjectMeta: metav1.ObjectMeta{
	  				Name: "dummy-host01",
	  			},
	  			Spec: v1.NodeSpec{
	  				Taints: []v1.Taint{},
	  			},
	  			Status: v1.NodeStatus{
	  				Conditions: []v1.NodeCondition{
	  					{
	  						Type:   v1.NodeReady,
	  						Status: v1.ConditionTrue,
	  					},
	  				},
	  			},
	  		},
	  		silent:        false,
	  		conditionType: v1.NodeReady,
	  		wantTrue:      true,
	  		expect:        true,
	  	},
	  	{
	  		desc: "Return should be false.  wantTrue: true, conditionType equls node's one, but node has taints. it's a first condition",
	  		node: &v1.Node{
	  			TypeMeta: metav1.TypeMeta{},
	  			ObjectMeta: metav1.ObjectMeta{
	  				Name: "dummy-host01",
	  			},
	  			Spec: v1.NodeSpec{
	  				Taints: []v1.Taint{
	  					{
	  						Key:    schedulerapi.TaintNodeUnreachable,
	  						Effect: v1.TaintEffectNoExecute,
	  					},
	  				},
	  			},
	  			Status: v1.NodeStatus{
	  				Conditions: []v1.NodeCondition{
	  					{
	  						Type:   v1.NodeReady,
	  						Status: v1.ConditionTrue,
	  					},
	  				},
	  			},
	  		},
	  		silent:        false,
	  		conditionType: v1.NodeReady,
	  		wantTrue:      true,
	  		expect:        false,
	  	},
	  	{
	  		desc: "Return should be false.  wantTrue: true, conditionType equls node's one, but node has taints it's a second condition",
	  		node: &v1.Node{
	  			TypeMeta: metav1.TypeMeta{},
	  			ObjectMeta: metav1.ObjectMeta{
	  				Name: "dummy-host01",
	  			},
	  			Spec: v1.NodeSpec{
	  				Taints: []v1.Taint{
	  					{
	  						Key:    schedulerapi.TaintNodeNotReady,
	  						Effect: v1.TaintEffectNoExecute,
	  					},
	  				},
	  			},
	  			Status: v1.NodeStatus{
	  				Conditions: []v1.NodeCondition{
	  					{
	  						Type:   v1.NodeReady,
	  						Status: v1.ConditionTrue,
	  					},
	  				},
	  			},
	  		},
	  		silent:        false,
	  		conditionType: v1.NodeReady,
	  		wantTrue:      true,
	  		expect:        false,
	  	},
	  	{
	  		desc: "Return should be true.  wantTrue: false, conditionType equls node's one, but status don't equal to ConditionTrue",
	  		node: &v1.Node{
	  			TypeMeta: metav1.TypeMeta{},
	  			ObjectMeta: metav1.ObjectMeta{
	  				Name: "dummy-host01",
	  			},
	  			Spec: v1.NodeSpec{
	  				Taints: []v1.Taint{
	  					{
	  						Key:    schedulerapi.TaintNodeUnreachable,
	  						Effect: v1.TaintEffectNoExecute,
	  					},
	  				},
	  			},
	  			Status: v1.NodeStatus{
	  				Conditions: []v1.NodeCondition{
	  					{
	  						Type:   v1.NodeReady,
	  						Status: v1.ConditionFalse,
	  					},
	  				},
	  			},
	  		},
	  		silent:        false,
	  		conditionType: v1.NodeReady,
	  		wantTrue:      false,
	  		expect:        true,
	  	},
	  	{ //***
	  		desc: "Return should be false.  wantTrue: false, conditionType equls node's one, and status equal to ConditionTrue",
	  		node: &v1.Node{
	  			TypeMeta: metav1.TypeMeta{},
	  			ObjectMeta: metav1.ObjectMeta{
	  				Name: "dummy-host01",
	  			},
	  			Spec: v1.NodeSpec{
	  				Taints: []v1.Taint{
	  					{
	  						Key:    schedulerapi.TaintNodeUnreachable,
	  						Effect: v1.TaintEffectNoExecute,
	  					},
	  				},
	  			},
	  			Status: v1.NodeStatus{
	  				Conditions: []v1.NodeCondition{
	  					{
	  						Type:   v1.NodeReady,
	  						Status: v1.ConditionTrue,
	  					},
	  				},
	  			},
	  		},
	  		silent:        false,
	  		conditionType: v1.NodeReady,
	  		wantTrue:      false,
	  		expect:        false,
	  	},
	  	{
	  		desc: "Return should be true.  wantTrue: true, conditionType equls to node's one, but don't equal to NodeReady, and status equal to ConditionTrue",
	  		node: &v1.Node{
	  			TypeMeta: metav1.TypeMeta{},
	  			ObjectMeta: metav1.ObjectMeta{
	  				Name: "dummy-host01",
	  			},
	  			Spec: v1.NodeSpec{
	  				Taints: []v1.Taint{
	  					{
	  						Key:    schedulerapi.TaintNodeUnreachable,
	  						Effect: v1.TaintEffectNoExecute,
	  					},
	  				},
	  			},
	  			Status: v1.NodeStatus{
	  				Conditions: []v1.NodeCondition{
	  					{
	  						Type:   v1.NodeMemoryPressure,
	  						Status: v1.ConditionTrue,
	  					},
	  				},
	  			},
	  		},
	  		silent:        false,
	  		conditionType: v1.NodeMemoryPressure,
	  		wantTrue:      true,
	  		expect:        true,
	  	},
	  	{
	  		desc: "Return should be true.  wantTrue: false, conditionType equls to node's one, but don't equal to NodeReady, and status don't equal to ConditionTrue",
	  		node: &v1.Node{
	  			TypeMeta: metav1.TypeMeta{},
	  			ObjectMeta: metav1.ObjectMeta{
	  				Name: "dummy-host01",
	  			},
	  			Spec: v1.NodeSpec{
	  				Taints: []v1.Taint{
	  					{
	  						Key:    schedulerapi.TaintNodeUnreachable,
	  						Effect: v1.TaintEffectNoExecute,
	  					},
	  				},
	  			},
	  			Status: v1.NodeStatus{
	  				Conditions: []v1.NodeCondition{
	  					{
	  						Type:   v1.NodeMemoryPressure,
	  						Status: v1.ConditionFalse,
	  					},
	  				},
	  			},
	  		},
	  		silent:        false,
	  		conditionType: v1.NodeMemoryPressure,
	  		wantTrue:      false,
	  		expect:        true,
	  	},
	  	{
	  		desc: "Return should be false.  wantTrue: true, conditionType equls to node's one, but don't equal to NodeReady, and status don't equal to ConditionTrue",
	  		node: &v1.Node{
	  			TypeMeta: metav1.TypeMeta{},
	  			ObjectMeta: metav1.ObjectMeta{
	  				Name: "dummy-host01",
	  			},
	  			Spec: v1.NodeSpec{
	  				Taints: []v1.Taint{
	  					{
	  						Key:    schedulerapi.TaintNodeUnreachable,
	  						Effect: v1.TaintEffectNoExecute,
	  					},
	  				},
	  			},
	  			Status: v1.NodeStatus{
	  				Conditions: []v1.NodeCondition{
	  					{
	  						Type:   v1.NodeMemoryPressure,
	  						Status: v1.ConditionFalse,
	  					},
	  				},
	  			},
	  		},
	  		silent:        false,
	  		conditionType: v1.NodeMemoryPressure,
	  		wantTrue:      true,
	  		expect:        false,
	  	},
	  	{
	  		desc: "Return should be false.  wantTrue: true, conditionType don't equl to node's one",
	  		node: &v1.Node{
	  			TypeMeta: metav1.TypeMeta{},
	  			ObjectMeta: metav1.ObjectMeta{
	  				Name: "dummy-host01",
	  			},
	  			Spec: v1.NodeSpec{
	  				Taints: []v1.Taint{
	  					{
	  						Key:    schedulerapi.TaintNodeUnreachable,
	  						Effect: v1.TaintEffectNoExecute,
	  					},
	  				},
	  			},
	  			Status: v1.NodeStatus{
	  				Conditions: []v1.NodeCondition{
	  					{
	  						Type:   v1.NodeMemoryPressure,
	  						Status: v1.ConditionFalse,
	  					},
	  				},
	  			},
	  		},
	  		silent:        false,
	  		conditionType: v1.NodeReady,
	  		wantTrue:      true,
	  		expect:        false,
	  	},
	  }*/
	Logf("Revised------------------------")

	for i, tc := range testCases {
		Logf("testcase:[%v]", i)
		result := isNodeConditionSetAsExpected5(tc.node, tc.conditionType, tc.wantTrue, tc.silent)
		if tc.expect != result {
			t.Errorf("testcase:[%v] error occurs: %v", i, tc.desc)
		}
	}

}

func setup() {
	testCases = []testCase{
		{
			desc: "Return should be true.  wantTrue: true, conditionType equls node's one, and node has no taints",
			node: &v1.Node{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name: "dummy-host01",
				},
				Spec: v1.NodeSpec{
					Taints: []v1.Taint{},
				},
				Status: v1.NodeStatus{
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeReady,
							Status: v1.ConditionTrue,
						},
					},
				},
			},
			silent:        false,
			conditionType: v1.NodeReady,
			wantTrue:      true,
			expect:        true,
		},
		{
			desc: "Return should be false.  wantTrue: true, conditionType equls node's one, but node has taints. it's a first condition",
			node: &v1.Node{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name: "dummy-host01",
				},
				Spec: v1.NodeSpec{
					Taints: []v1.Taint{
						{
							Key:    schedulerapi.TaintNodeUnreachable,
							Effect: v1.TaintEffectNoExecute,
						},
					},
				},
				Status: v1.NodeStatus{
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeReady,
							Status: v1.ConditionTrue,
						},
					},
				},
			},
			silent:        false,
			conditionType: v1.NodeReady,
			wantTrue:      true,
			expect:        false,
		},
		{
			desc: "Return should be false.  wantTrue: true, conditionType equls node's one, but node has taints it's a second condition",
			node: &v1.Node{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name: "dummy-host01",
				},
				Spec: v1.NodeSpec{
					Taints: []v1.Taint{
						{
							Key:    schedulerapi.TaintNodeNotReady,
							Effect: v1.TaintEffectNoExecute,
						},
					},
				},
				Status: v1.NodeStatus{
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeReady,
							Status: v1.ConditionTrue,
						},
					},
				},
			},
			silent:        false,
			conditionType: v1.NodeReady,
			wantTrue:      true,
			expect:        false,
		},
		{
			desc: "Return should be true.  wantTrue: false, conditionType equls node's one, but status don't equal to ConditionTrue",
			node: &v1.Node{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name: "dummy-host01",
				},
				Spec: v1.NodeSpec{
					Taints: []v1.Taint{
						{
							Key:    schedulerapi.TaintNodeUnreachable,
							Effect: v1.TaintEffectNoExecute,
						},
					},
				},
				Status: v1.NodeStatus{
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeReady,
							Status: v1.ConditionFalse,
						},
					},
				},
			},
			silent:        false,
			conditionType: v1.NodeReady,
			wantTrue:      false,
			expect:        true,
		},
		{ //***
			desc: "Return should be false.  wantTrue: false, conditionType equls node's one, and status equal to ConditionTrue",
			node: &v1.Node{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name: "dummy-host01",
				},
				Spec: v1.NodeSpec{
					Taints: []v1.Taint{
						{
							Key:    schedulerapi.TaintNodeUnreachable,
							Effect: v1.TaintEffectNoExecute,
						},
					},
				},
				Status: v1.NodeStatus{
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeReady,
							Status: v1.ConditionTrue,
						},
					},
				},
			},
			silent:        false,
			conditionType: v1.NodeReady,
			wantTrue:      false,
			expect:        false,
		},
		{
			desc: "Return should be true.  wantTrue: true, conditionType equls to node's one, but don't equal to NodeReady, and status equal to ConditionTrue",
			node: &v1.Node{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name: "dummy-host01",
				},
				Spec: v1.NodeSpec{
					Taints: []v1.Taint{
						{
							Key:    schedulerapi.TaintNodeUnreachable,
							Effect: v1.TaintEffectNoExecute,
						},
					},
				},
				Status: v1.NodeStatus{
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeMemoryPressure,
							Status: v1.ConditionTrue,
						},
					},
				},
			},
			silent:        false,
			conditionType: v1.NodeMemoryPressure,
			wantTrue:      true,
			expect:        true,
		},
		{
			desc: "Return should be true.  wantTrue: false, conditionType equls to node's one, but don't equal to NodeReady, and status don't equal to ConditionTrue",
			node: &v1.Node{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name: "dummy-host01",
				},
				Spec: v1.NodeSpec{
					Taints: []v1.Taint{
						{
							Key:    schedulerapi.TaintNodeUnreachable,
							Effect: v1.TaintEffectNoExecute,
						},
					},
				},
				Status: v1.NodeStatus{
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeMemoryPressure,
							Status: v1.ConditionFalse,
						},
					},
				},
			},
			silent:        false,
			conditionType: v1.NodeMemoryPressure,
			wantTrue:      false,
			expect:        true,
		},
		{
			desc: "Return should be false.  wantTrue: true, conditionType equls to node's one, but don't equal to NodeReady, and status don't equal to ConditionTrue",
			node: &v1.Node{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name: "dummy-host01",
				},
				Spec: v1.NodeSpec{
					Taints: []v1.Taint{
						{
							Key:    schedulerapi.TaintNodeUnreachable,
							Effect: v1.TaintEffectNoExecute,
						},
					},
				},
				Status: v1.NodeStatus{
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeMemoryPressure,
							Status: v1.ConditionFalse,
						},
					},
				},
			},
			silent:        false,
			conditionType: v1.NodeMemoryPressure,
			wantTrue:      true,
			expect:        false,
		},
		{
			desc: "Return should be false.  wantTrue: true, conditionType don't equl to node's one",
			node: &v1.Node{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name: "dummy-host01",
				},
				Spec: v1.NodeSpec{
					Taints: []v1.Taint{
						{
							Key:    schedulerapi.TaintNodeUnreachable,
							Effect: v1.TaintEffectNoExecute,
						},
					},
				},
				Status: v1.NodeStatus{
					Conditions: []v1.NodeCondition{
						{
							Type:   v1.NodeMemoryPressure,
							Status: v1.ConditionFalse,
						},
					},
				},
			},
			silent:        false,
			conditionType: v1.NodeReady,
			wantTrue:      true,
			expect:        false,
		},
	}

}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	//teardown()
	os.Exit(ret)
}
