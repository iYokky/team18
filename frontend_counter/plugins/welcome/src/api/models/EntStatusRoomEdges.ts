/* tslint:disable */
/* eslint-disable */
/**
 * SUT SE Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDataRoom,
    EntDataRoomFromJSON,
    EntDataRoomFromJSONTyped,
    EntDataRoomToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStatusRoomEdges
 */
export interface EntStatusRoomEdges {
    /**
     * Datarooms holds the value of the datarooms edge.
     * @type {Array<EntDataRoom>}
     * @memberof EntStatusRoomEdges
     */
    datarooms?: Array<EntDataRoom>;
}

export function EntStatusRoomEdgesFromJSON(json: any): EntStatusRoomEdges {
    return EntStatusRoomEdgesFromJSONTyped(json, false);
}

export function EntStatusRoomEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStatusRoomEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'datarooms': !exists(json, 'datarooms') ? undefined : ((json['datarooms'] as Array<any>).map(EntDataRoomFromJSON)),
    };
}

export function EntStatusRoomEdgesToJSON(value?: EntStatusRoomEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'datarooms': value.datarooms === undefined ? undefined : ((value.datarooms as Array<any>).map(EntDataRoomToJSON)),
    };
}


