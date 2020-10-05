import { doFetch } from '../client';
import ActionTypes from '../action_types';
import {id as pluginId} from '../manifest';
import {getPluginServerRoute} from '../selectors';

export function fetchUserReputations(userId) {
    return async (dispatch, getState) => {
        const baseUrl = getPluginServerRoute(getState());
        let data = null;
        try {
            data = await doFetch(`${baseUrl}/api/v1/reputation?user_id=${userId}`, {
                method: 'get'
            });
            dispatch({
                type: ActionTypes.RECEIVED_USER_REPUTATIONS,
                userId,
                data: data
            });
    
        }
        catch (error) {
            return {error};
        }

        if (data.error) {
            return {error: new Error(data.error)};
        }

        return {data};
    }
}