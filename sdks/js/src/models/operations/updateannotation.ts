/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import * as components from "../components";
import * as z from "zod";

export type UpdateAnnotationRequest = {
    /**
     * The id of the annotation.
     */
    id: string;
    /**
     * The fields to update.
     */
    baseAnnotation: components.BaseAnnotation;
};

export type UpdateAnnotationResponse = {
    httpMeta: components.HTTPMetadata;
    /**
     * The updated annotation.
     */
    annotation?: components.Annotation | undefined;
};

/** @internal */
export namespace UpdateAnnotationRequest$ {
    export type Inbound = {
        id: string;
        BaseAnnotation: components.BaseAnnotation$.Inbound;
    };

    export const inboundSchema: z.ZodType<UpdateAnnotationRequest, z.ZodTypeDef, Inbound> = z
        .object({
            id: z.string(),
            BaseAnnotation: components.BaseAnnotation$.inboundSchema,
        })
        .transform((v) => {
            return {
                id: v.id,
                baseAnnotation: v.BaseAnnotation,
            };
        });

    export type Outbound = {
        id: string;
        BaseAnnotation: components.BaseAnnotation$.Outbound;
    };

    export const outboundSchema: z.ZodType<Outbound, z.ZodTypeDef, UpdateAnnotationRequest> = z
        .object({
            id: z.string(),
            baseAnnotation: components.BaseAnnotation$.outboundSchema,
        })
        .transform((v) => {
            return {
                id: v.id,
                BaseAnnotation: v.baseAnnotation,
            };
        });
}

/** @internal */
export namespace UpdateAnnotationResponse$ {
    export type Inbound = {
        HttpMeta: components.HTTPMetadata$.Inbound;
        Annotation?: components.Annotation$.Inbound | undefined;
    };

    export const inboundSchema: z.ZodType<UpdateAnnotationResponse, z.ZodTypeDef, Inbound> = z
        .object({
            HttpMeta: components.HTTPMetadata$.inboundSchema,
            Annotation: components.Annotation$.inboundSchema.optional(),
        })
        .transform((v) => {
            return {
                httpMeta: v.HttpMeta,
                ...(v.Annotation === undefined ? null : { annotation: v.Annotation }),
            };
        });

    export type Outbound = {
        HttpMeta: components.HTTPMetadata$.Outbound;
        Annotation?: components.Annotation$.Outbound | undefined;
    };

    export const outboundSchema: z.ZodType<Outbound, z.ZodTypeDef, UpdateAnnotationResponse> = z
        .object({
            httpMeta: components.HTTPMetadata$.outboundSchema,
            annotation: components.Annotation$.outboundSchema.optional(),
        })
        .transform((v) => {
            return {
                HttpMeta: v.httpMeta,
                ...(v.annotation === undefined ? null : { Annotation: v.annotation }),
            };
        });
}
