/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import * as components from "../components";
import * as z from "zod";

export type DeleteAnnotationRequest = {
    /**
     * The id of the annotation.
     */
    id: string;
};

export type DeleteAnnotationResponse = {
    httpMeta: components.HTTPMetadata;
};

/** @internal */
export namespace DeleteAnnotationRequest$ {
    export type Inbound = {
        id: string;
    };

    export const inboundSchema: z.ZodType<DeleteAnnotationRequest, z.ZodTypeDef, Inbound> = z
        .object({
            id: z.string(),
        })
        .transform((v) => {
            return {
                id: v.id,
            };
        });

    export type Outbound = {
        id: string;
    };

    export const outboundSchema: z.ZodType<Outbound, z.ZodTypeDef, DeleteAnnotationRequest> = z
        .object({
            id: z.string(),
        })
        .transform((v) => {
            return {
                id: v.id,
            };
        });
}

/** @internal */
export namespace DeleteAnnotationResponse$ {
    export type Inbound = {
        HttpMeta: components.HTTPMetadata$.Inbound;
    };

    export const inboundSchema: z.ZodType<DeleteAnnotationResponse, z.ZodTypeDef, Inbound> = z
        .object({
            HttpMeta: components.HTTPMetadata$.inboundSchema,
        })
        .transform((v) => {
            return {
                httpMeta: v.HttpMeta,
            };
        });

    export type Outbound = {
        HttpMeta: components.HTTPMetadata$.Outbound;
    };

    export const outboundSchema: z.ZodType<Outbound, z.ZodTypeDef, DeleteAnnotationResponse> = z
        .object({
            httpMeta: components.HTTPMetadata$.outboundSchema,
        })
        .transform((v) => {
            return {
                HttpMeta: v.httpMeta,
            };
        });
}
