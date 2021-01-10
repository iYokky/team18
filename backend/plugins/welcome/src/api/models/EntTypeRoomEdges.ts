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
 * @interface EntTypeRoomEdges
 */
export interface EntTypeRoomEdges {
    /**
     * Datarooms holds the value of the datarooms edge.
     * @type {Array<EntDataRoom>}
     * @memberof EntTypeRoomEdges
     */
    datarooms?: Array<EntDataRoom>;
}

export function EntTypeRoomEdgesFromJSON(json: any): EntTypeRoomEdges {
    return EntTypeRoomEdgesFromJSONTyped(json, false);
}

export function EntTypeRoomEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTypeRoomEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'datarooms': !exists(json, 'datarooms') ? undefined : ((json['datarooms'] as Array<any>).map(EntDataRoomFromJSON)),
    };
}

export function EntTypeRoomEdgesToJSON(value?: EntTypeRoomEdges | null): any {
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

