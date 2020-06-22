/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by injection-gen. DO NOT EDIT.

package ingress

import (
	context "context"

	v1 "k8s.io/api/core/v1"
	v1alpha1 "knative.dev/networking/pkg/apis/networking/v1alpha1"
	ingress "knative.dev/networking/pkg/client/injection/reconciler/networking/v1alpha1/ingress"
	reconciler "knative.dev/pkg/reconciler"
)

// TODO: PLEASE COPY AND MODIFY THIS FILE AS A STARTING POINT

// newReconciledNormal makes a new reconciler event with event type Normal, and
// reason IngressReconciled.
func newReconciledNormal(namespace, name string) reconciler.Event {
	return reconciler.NewEvent(v1.EventTypeNormal, "IngressReconciled", "Ingress reconciled: \"%s/%s\"", namespace, name)
}

// Reconciler implements controller.Reconciler for Ingress resources.
type Reconciler struct {
	// TODO: add additional requirements here.
}

// Check that our Reconciler implements Interface
var _ ingress.Interface = (*Reconciler)(nil)

// Optionally check that our Reconciler implements Finalizer
//var _ ingress.Finalizer = (*Reconciler)(nil)

// Optionally check that our Reconciler implements ReadOnlyInterface
// Implement this to observe resources even when we are not the leader.
//var _ ingress.ReadOnlyInterface = (*Reconciler)(nil)

// Optionally check that our Reconciler implements ReadOnlyFinalizer
// Implement this to observe tombstoned resources even when we are not
// the leader (best effort).
//var _ ingress.ReadOnlyFinalizer = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, o *v1alpha1.Ingress) reconciler.Event {
	// TODO: use this if the resource implements InitializeConditions.
	// o.Status.InitializeConditions()

	// TODO: add custom reconciliation logic here.

	// TODO: use this if the object has .status.ObservedGeneration.
	// o.Status.ObservedGeneration = o.Generation
	return newReconciledNormal(o.Namespace, o.Name)
}

// Optionally, use FinalizeKind to add finalizers. FinalizeKind will be called
// when the resource is deleted.
//func (r *Reconciler) FinalizeKind(ctx context.Context, o *v1alpha1.Ingress) reconciler.Event {
//	// TODO: add custom finalization logic here.
//	return nil
//}

// Optionally, use ObserveKind to observe the resource when we are not the leader.
// func (r *Reconciler) ObserveKind(ctx context.Context, o *v1alpha1.Ingress) reconciler.Event {
// 	// TODO: add custom observation logic here.
// 	return nil
// }

// Optionally, use ObserveFinalizeKind to observe resources being finalized when we are no the leader.
//func (r *Reconciler) ObserveFinalizeKind(ctx context.Context, o *v1alpha1.Ingress) reconciler.Event {
// 	// TODO: add custom observation logic here.
//	return nil
//}
